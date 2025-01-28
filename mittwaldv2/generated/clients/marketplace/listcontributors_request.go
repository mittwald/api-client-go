package marketplace

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListContributorsRequest models a request for the 'extension-list-contributors'
// operation. See [1] for more information.
//
// List Contributors.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-list-contributors
type ListContributorsRequest struct {
	Limit *int64
	Skip  *int64
	Page  *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListContributorsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListContributorsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListContributorsRequest) url() string {
	return "/v2/contributors"
}

func (r *ListContributorsRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	return q
}
