package httpclient

import "net/http"

type StubClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *StubClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
