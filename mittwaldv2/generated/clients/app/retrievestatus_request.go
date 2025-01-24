package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type RetrieveStatusRequest struct {
	AppInstallationID string
}

func (r *RetrieveStatusRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *RetrieveStatusRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *RetrieveStatusRequest) url() string {
	return fmt.Sprintf("/v2/app-installations/%s/status", url.PathEscape(r.AppInstallationID))
}

func (r *RetrieveStatusRequest) query() url.Values {
	return nil
}
