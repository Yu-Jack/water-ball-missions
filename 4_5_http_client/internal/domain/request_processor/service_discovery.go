package request_processor

import (
	"errors"
	"strings"
	"sync"
	"time"

	"http.client/internal/domain"
)

// for demo
var data = "example.com 1.1.1.1 2.2.2.2 3.3.3.3 4.4.4.4 5.5.5.5" // fake, hard code testing

type serviceDiscovery struct {
	processor
	mapping   map[string][]string
	ipCounter map[string]int
	ipFailed  map[string]struct{}
	lock      sync.Mutex
}

func NewServiceDiscovery(next domain.RequestProcessor) domain.RequestProcessor {
	s := &serviceDiscovery{
		processor: processor{next: next},
		mapping:   make(map[string][]string),
		ipCounter: make(map[string]int),
		ipFailed:  make(map[string]struct{}),
		lock:      sync.Mutex{},
	}

	datas := strings.Split(data, " ")
	var ips []string
	for i := 1; i < len(datas); i++ {
		ips = append(ips, datas[i])
	}

	s.mapping[datas[0]] = ips
	s.ipCounter[datas[0]] = 0

	return s
}

func (s *serviceDiscovery) SendRequest(request domain.HttpRequest) error {
	if ips, ok := s.mapping[request.OriginalHost]; !ok {
		return errors.New("fail to resolve host")
	} else {
		s.findAvailableIP(ips, &request)

		err := s.processor.SendRequest(request)

		if err != nil {
			s.unreachableHandler(&request)
			return s.SendRequest(request) // 重 call 直到找到可用的 IP
		}

		return nil
	}
}

func (s *serviceDiscovery) unreachableHandler(request *domain.HttpRequest) {
	s.ipFailed[request.Hosts[0]] = struct{}{}
	s.ipCounter[request.OriginalHost]++ // 找不到指標重來

	time.AfterFunc(2*time.Second, func() { // 改成兩秒後變回正常，比較好測試
		s.lock.Lock()
		delete(s.ipFailed, request.Hosts[0]) // release IP
		s.lock.Unlock()
	})
}

func (s *serviceDiscovery) findAvailableIP(ips []string, request *domain.HttpRequest) {
	request.AvailableIPs = []string{}

	s.ipCounter[request.OriginalHost] = 0 // start from first one

	for {
		if _, ok := s.ipFailed[ips[s.ipCounter[request.OriginalHost]]]; ok {
			s.ipCounter[request.OriginalHost]++
			continue
		}

		request.Hosts = []string{
			ips[s.ipCounter[request.OriginalHost]%len(ips)],
		}

		break
	}

	for _, ip := range ips {
		if _, ok := s.ipFailed[ip]; !ok {
			request.AvailableIPs = append(request.AvailableIPs, ip)
		}
	}
}
