package domain

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListDomainsRequest models a request for the 'domain-list-domains' operation. See
// [1] for more information.
//
// # List Domains
//
// [1]: https://developer.mittwald.de/docs/v2/reference/domain/domain-list-domains
type ListDomainsRequest struct {
	ProjectID        *string
	Page             *int64
	Limit            *int64
	DomainSearchName *string
	ContactHash      *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListDomainsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListDomainsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListDomainsRequest) url() string {
	return "/v2/domains"
}

func (r *ListDomainsRequest) query() url.Values {
	q := make(url.Values)
	if r.ProjectID != nil {
		q.Set("projectId", *r.ProjectID)
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.DomainSearchName != nil {
		q.Set("domainSearchName", *r.DomainSearchName)
	}
	if r.ContactHash != nil {
		q.Set("contactHash", *r.ContactHash)
	}
	return q
}
