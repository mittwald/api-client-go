package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListRegistriesRequest models a request for the 'container-list-registries'
// operation. See [1] for more information.
//
// List Registries belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-list-registries
type ListRegistriesRequest struct {
	ProjectID      string
	HasCredentials *bool
	Limit          *int64
	Skip           *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListRegistriesRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListRegistriesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListRegistriesRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/registries", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListRegistriesRequest) query() url.Values {
	q := make(url.Values)
	if r.HasCredentials != nil {
		q.Set("hasCredentials", strconv.FormatBool(*r.HasCredentials))
	}
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
