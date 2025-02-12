package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// StartServiceRequest models a request for the 'container-start-service'
// operation. See [1] for more information.
//
// Start a stopped Service.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-start-service
type StartServiceRequest struct {
	StackID   string
	ServiceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *StartServiceRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *StartServiceRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *StartServiceRequest) url() string {
	return fmt.Sprintf("/v2/stacks/%s/services/%s/actions/start", url.PathEscape(r.StackID), url.PathEscape(r.ServiceID))
}

func (r *StartServiceRequest) query() url.Values {
	return nil
}
