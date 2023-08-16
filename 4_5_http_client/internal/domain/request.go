package domain

type HttpRequest struct {
	Scheme       string
	Hosts        []string
	AvailableIPs []string
	Path         string
	Method       string
	OriginalHost string
}

func NewHttpRequest(host, path string) HttpRequest {
	return HttpRequest{
		Scheme:       "http",
		Hosts:        []string{host}, // 實際會用這個 call request
		AvailableIPs: []string{},     // 遇到服務探索才會需要 assign 進來
		Path:         path,
		Method:       "GET",
		OriginalHost: host,
	}
}
