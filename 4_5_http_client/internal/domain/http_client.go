package domain

type RequestProcessor interface {
	SendRequest(request HttpRequest) error
}

type HttpClient interface {
	SendRequest(request HttpRequest) error
}
