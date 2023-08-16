package domain

type RequestProcessor interface {
	SendRequest(request HttpRequest) (HttpRequest, error)
}

type HttpClient interface {
	SendRequest(request HttpRequest) (HttpRequest, error)
}
