package test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"http.client/internal/domain"
	"http.client/internal/domain/request_processor"
)

func TestHttpClientSuite(t *testing.T) {
	suite.Run(t, new(httpClientSuite))
}

type httpClientSuite struct {
	suite.Suite
	request domain.HttpRequest
}

func (s *httpClientSuite) SetupSuite() {
	s.request = domain.NewHttpRequest("example.com", "/hi")
}

func (s *httpClientSuite) captureConsole(cb func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cb()

	_ = w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

func (s *httpClientSuite) Test_Nothing() {
	client := domain.NewFakeHttpClient()

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://example.com/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_Blacklist() {
	client := request_processor.NewBlacklist(
		domain.NewFakeHttpClient(),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://example.com/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_Discovery() {
	client := request_processor.NewServiceDiscovery(
		domain.NewFakeHttpClient(),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
		time.Sleep(3 * time.Second)
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://2.2.2.2/hi\n[SUCCESS] http://1.1.1.1/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_ServiceDiscovery_Blacklist() {
	client := request_processor.NewServiceDiscovery(
		request_processor.NewBlacklist(
			domain.NewFakeHttpClient(),
		),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://3.3.3.3/hi\n[SUCCESS] http://3.3.3.3/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_ServiceDiscovery_LoadBalancer() {
	client := request_processor.NewServiceDiscovery(
		request_processor.NewLoadBalancer(
			domain.NewFakeHttpClient(),
		),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://3.3.3.3/hi\n[SUCCESS] http://4.4.4.4/hi\n[SUCCESS] http://5.5.5.5/hi\n[SUCCESS] http://2.2.2.2/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_ServiceDiscovery_LoadBalancer_Blacklist() {
	client := request_processor.NewServiceDiscovery(
		request_processor.NewLoadBalancer(
			request_processor.NewBlacklist(
				domain.NewFakeHttpClient(),
			),
		),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://3.3.3.3/hi\n[SUCCESS] http://4.4.4.4/hi\n[SUCCESS] http://5.5.5.5/hi\n[SUCCESS] http://5.5.5.5/hi\n[SUCCESS] http://3.3.3.3/hi\n",
		output,
		"order should be right",
	)
}

func (s *httpClientSuite) Test_Blacklist_LoadBalancer_ServiceDiscovery() {
	client := request_processor.NewBlacklist(
		request_processor.NewLoadBalancer(
			request_processor.NewServiceDiscovery(
				domain.NewFakeHttpClient(),
			),
		),
	)

	output := s.captureConsole(func() {
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
		_ = client.SendRequest(s.request)
	})

	s.Require().Equal(
		"[SUCCESS] http://2.2.2.2/hi\n[SUCCESS] http://2.2.2.2/hi\n[SUCCESS] http://2.2.2.2/hi\n[SUCCESS] http://2.2.2.2/hi\n[SUCCESS] http://2.2.2.2/hi\n",
		output,
		"order should be right",
	)
}
