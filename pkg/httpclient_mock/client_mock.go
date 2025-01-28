package httpclient_mock

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// MockRequestRunner is a helper class that implements the httpclient.RequestRunner
// interface and may be used as a mock implementation during testing.
type MockRequestRunner struct {
	Requests []http.Request
	Matchers map[string]func(r *http.Request) *http.Response
}

// ExpectRequest configures the mock client to expect an HTTP request with a
// given method and path. The response that should be returned can be configured
// by providing a list of ResponseOption's.
func (m *MockRequestRunner) ExpectRequest(method, path string, opts ...ResponseOption) {
	bodyReader := strings.NewReader("")
	bodyReadCloser := io.NopCloser(bodyReader)

	resp := http.Response{Body: bodyReadCloser, StatusCode: 204, Status: http.StatusText(204)}

	for _, o := range opts {
		o(&resp)
	}

	m.ExpectRequestWithResponse(method, path, &resp)
}

func (m *MockRequestRunner) ExpectRequestWithResponse(method, path string, resp *http.Response) {
	m.ExpectRequestWithResponseFunc(method, path, func(*http.Request) *http.Response { return resp })
}

func (m *MockRequestRunner) ExpectRequestWithResponseFunc(method, path string, resp func(r *http.Request) *http.Response) {
	if m.Matchers == nil {
		m.Matchers = make(map[string]func(r *http.Request) *http.Response)
	}

	key := strings.ToLower(method + "_" + path)
	m.Matchers[key] = resp
}

func (m *MockRequestRunner) Do(request *http.Request) (*http.Response, error) {
	m.Requests = append(m.Requests, *request)

	key := strings.ToLower(request.Method + "_" + request.URL.Path)
	if handler, ok := m.Matchers[key]; ok {
		return handler(request), nil
	}

	return nil, fmt.Errorf("unexpected %s request to %s", request.Method, request.URL)
}
