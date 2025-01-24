package database

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListMysqlCharsetsRequest struct {
	Version *string
}

func (r *ListMysqlCharsetsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListMysqlCharsetsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListMysqlCharsetsRequest) url() string {
	return "/v2/mysql-charsets"
}

func (r *ListMysqlCharsetsRequest) query() url.Values {
	q := make(url.Values)
	if r.Version != nil {
		q.Set("version", *r.Version)
	}
	return q
}
