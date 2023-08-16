package domain

import (
	"errors"
	"fmt"
)

type fakeHttpClient struct {
	req   HttpRequest
	count int
}

func NewFakeHttpClient() HttpClient {
	return &fakeHttpClient{
		count: 0,
	}
}

func (f *fakeHttpClient) SendRequest(request HttpRequest) error {
	if len(request.Hosts) == 0 {
		return errors.New("fail to send request")
	}

	host := request.Hosts[0]

	if f.monitorSomeSituations(request) {
		return errors.New("fake data")
	}

	fmt.Printf("[SUCCESS] %s://%s%s\n", request.Scheme, host, request.Path)

	f.count++

	return nil
}

func (f *fakeHttpClient) monitorSomeSituations(request HttpRequest) bool {
	if request.Hosts[0] == "1.1.1.1" && f.count == 0 { // 模擬第一次打不到, hard code testing
		return true
	}

	return false
}
