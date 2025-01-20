package user

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListSessionsRequest struct {
}

func (r *ListSessionsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListSessionsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListSessionsRequest) url() string {
	return "/v2/users/self/sessions"
}

func (r *ListSessionsRequest) query() url.Values {
	return nil
}
