package customer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListInvitesForCustomerRequest struct {
	CustomerID string
	Limit      *int64
	Skip       *int64
}

func (r *ListInvitesForCustomerRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListInvitesForCustomerRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListInvitesForCustomerRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/invites", url.PathEscape(r.CustomerID))
}

func (r *ListInvitesForCustomerRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
