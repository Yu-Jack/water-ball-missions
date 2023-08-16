package request_processor

import "http.client/internal/domain"

type processor struct {
	next domain.RequestProcessor
}

func (p *processor) GetNext() domain.RequestProcessor {
	return p.next
}

func (p *processor) SendRequest(request domain.HttpRequest) error {
	if p.next != nil {
		return p.next.SendRequest(request)
	}
	return nil
}
