package request_processor

import (
	"fmt"

	"http.client/internal/domain"
)

type blacklist struct {
	processor
	lists map[string]struct{}
}

func NewBlacklist(next domain.RequestProcessor) domain.RequestProcessor {
	s := &blacklist{
		processor: processor{next: next},
		lists:     make(map[string]struct{}),
	}

	s.lists["2.2.2.2"] = struct{}{} // fake, hard code testing

	return s
}

func (s *blacklist) SendRequest(request domain.HttpRequest) error {
	if _, ok := s.lists[request.Hosts[0]]; ok {
		return fmt.Errorf("%s is in blacklist", request.Hosts[0])
	}

	return s.processor.SendRequest(request)
}
