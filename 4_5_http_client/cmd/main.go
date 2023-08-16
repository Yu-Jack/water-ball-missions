package main

import (
	"fmt"

	"http.client/internal/domain"
	"http.client/internal/domain/request_processor"
)

type CustomHttpClient struct{}

func (c *CustomHttpClient) SendRequest(request domain.HttpRequest) (domain.HttpRequest, error) {
	fmt.Printf("%#v\n", request)
	fmt.Println("Do nothing because I'm customer.")
	return request, nil
}

func main() {
	client := request_processor.NewServiceDiscovery(
		request_processor.NewLoadBalancer(
			request_processor.NewBlacklist(
				&CustomHttpClient{},
			),
		),
	)

	_, _ = client.SendRequest(domain.NewHttpRequest("example.com", "/hi"))
}
