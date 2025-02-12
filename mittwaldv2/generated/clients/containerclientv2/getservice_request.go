package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetServiceRequest models a request for the 'container-get-service' operation.
// See [1] for more information.
//
// Get a Service belonging to a Stack.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-get-service
type GetServiceRequest struct {
	StackID   string
	ServiceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetServiceRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *GetServiceRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetServiceRequest) url() string {
	return fmt.Sprintf("/v2/stacks/%s/services/%s", url.PathEscape(r.StackID), url.PathEscape(r.ServiceID))
}

func (r *GetServiceRequest) query() url.Values {
	return nil
}
