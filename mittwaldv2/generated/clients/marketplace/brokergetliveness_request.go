package marketplace

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type BrokerGetLivenessRequest struct {
}

func (r *BrokerGetLivenessRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *BrokerGetLivenessRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *BrokerGetLivenessRequest) url() string {
	return "/v2/broker/health/liveness"
}

func (r *BrokerGetLivenessRequest) query() url.Values {
	return nil
}
