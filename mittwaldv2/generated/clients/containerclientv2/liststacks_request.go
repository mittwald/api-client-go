package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListStacksRequest models a request for the 'container-list-stacks' operation.
// See [1] for more information.
//
// List Stacks belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-list-stacks
type ListStacksRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListStacksRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListStacksRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListStacksRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/stacks", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *ListStacksRequest) query() url.Values {
	return nil
}
