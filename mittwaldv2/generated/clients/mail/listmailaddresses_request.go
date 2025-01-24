package mail

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListMailAddressesRequest struct {
	ProjectID string
	Search    *string
	Limit     *int64
	Skip      *int64
	Page      *int64
}

func (r *ListMailAddressesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListMailAddressesRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListMailAddressesRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/mail-addresses", url.PathEscape(r.ProjectID))
}

func (r *ListMailAddressesRequest) query() url.Values {
	q := make(url.Values)
	if r.Search != nil {
		q.Set("search", *r.Search)
	}
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
