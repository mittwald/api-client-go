package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListInvitesForCustomerRequest models a request for the
// 'customer-list-invites-for-customer' operation. See [1] for more information.
//
// List Invites belonging to a Customer.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-list-invites-for-customer
type ListInvitesForCustomerRequest struct {
	CustomerID string
	Limit      *int64
	Skip       *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListInvitesForCustomerRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListInvitesForCustomerRequest) body() (io.Reader, string, error) {
	return nil, "", nil
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
