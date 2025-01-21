package domain

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeprecatedGetHandleFieldsRequest struct {
	DomainName string
}

func (r *DeprecatedGetHandleFieldsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *DeprecatedGetHandleFieldsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedGetHandleFieldsRequest) url() string {
	return fmt.Sprintf("/v2/domains/handle-schema/%s", url.PathEscape(r.DomainName))
}

func (r *DeprecatedGetHandleFieldsRequest) query() url.Values {
	return nil
}