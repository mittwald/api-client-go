package httpclient

import (
	"net/http"
)

type authenticatedClient struct {
	apiToken string
	inner    RequestRunner
}

func NewAuthenticatedClient(inner RequestRunner, apiToken string) RequestRunner {
	return &authenticatedClient{apiToken, inner}
}

func (a *authenticatedClient) Do(request *http.Request) (*http.Response, error) {
	request.Header.Set("X-Access-Token", a.apiToken)
	return a.inner.Do(request)
}
