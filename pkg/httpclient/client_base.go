package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type baseUrlClient struct {
	inner   RequestRunner
	baseURL *url.URL
}

func NewClientWithBaseURL(inner RequestRunner, baseURL string) (RequestRunner, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL '%s': %w", baseURL, err)
	}

	return &baseUrlClient{
		inner:   inner,
		baseURL: u,
	}, nil
}

func (b *baseUrlClient) Do(request *http.Request) (*http.Response, error) {
	request.URL.Scheme = b.baseURL.Scheme
	request.URL.Host = b.baseURL.Host

	if !strings.HasPrefix(request.URL.Path, "/") {
		request.URL.Path = strings.TrimSuffix(b.baseURL.Path, "/") + "/" + request.URL.Path
	}

	return b.inner.Do(request)
}
