package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// RestartServiceRequest models a request for the 'container-restart-service'
// operation. See [1] for more information.
//
// Restart a started Service.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-restart-service
type RestartServiceRequest struct {
	StackID   string
	ServiceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *RestartServiceRequest) BuildRequest() (*http.Request, error) {
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

func (r *RestartServiceRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *RestartServiceRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/stacks/%s/services/%s/actions/restart", url.PathEscape(r.StackID), url.PathEscape(r.ServiceID)),
	}
	return u.String()
}

func (r *RestartServiceRequest) query() url.Values {
	return nil
}
