package request_processor

import (
	"http.client/internal/domain"
)

type loadBalancer struct {
	processor
	ipCounter map[string]int
}

func NewLoadBalancer(next domain.RequestProcessor) domain.RequestProcessor {
	s := &loadBalancer{
		processor: processor{next: next},
		ipCounter: make(map[string]int),
	}

	return s
}

func (l *loadBalancer) SendRequest(request domain.HttpRequest) error {

	if len(request.AvailableIPs) != 0 {
		if _, ok := l.ipCounter[request.OriginalHost]; !ok {
			l.ipCounter[request.OriginalHost] = 0
		} else {
			l.ipCounter[request.OriginalHost]++
		}

		request.Hosts = []string{
			request.AvailableIPs[l.ipCounter[request.OriginalHost]%len(request.AvailableIPs)],
		}
	}

	return l.processor.SendRequest(request)
}
