package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

type RefreshFunc func() (string, time.Time, error)

type authenticatedClient struct {
	apiToken           string
	apiTokenHeaderName string
	apiTokenExpiration *time.Time
	refreshFunc        RefreshFunc
	inner              RequestRunner
}

func NewAuthenticatedClient(inner RequestRunner, apiToken, apiTokenHeaderName string) RequestRunner {
	return &authenticatedClient{
		apiToken:           apiToken,
		apiTokenHeaderName: apiTokenHeaderName,
		inner:              inner,
	}
}

func NewAuthenticatedClientWithRefresh(inner RequestRunner, apiToken, apiTokenHeaderName string, expiration time.Time, refreshFunc RefreshFunc) RequestRunner {
	return &authenticatedClient{
		apiToken:           apiToken,
		apiTokenHeaderName: apiTokenHeaderName,
		apiTokenExpiration: &expiration,
		refreshFunc:        refreshFunc,
		inner:              inner,
	}
}

func (a *authenticatedClient) Do(request *http.Request) (*http.Response, error) {
	if err := a.refreshTokenIfNecessary(); err != nil {
		return nil, err
	}

	request.Header.Set(a.apiTokenHeaderName, a.apiToken)
	return a.inner.Do(request)
}

func (a *authenticatedClient) refreshTokenIfNecessary() error {
	if a.apiTokenExpiration == nil || a.refreshFunc == nil {
		return nil
	}

	if a.apiTokenExpiration.Before(time.Now()) {
		newToken, newExpiration, err := a.refreshFunc()
		if err != nil {
			return fmt.Errorf("error while refreshing API token: %w", err)
		}

		a.apiToken = newToken
		a.apiTokenExpiration = &newExpiration
	}

	return nil
}
