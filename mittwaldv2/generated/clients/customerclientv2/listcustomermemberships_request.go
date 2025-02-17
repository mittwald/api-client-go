package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListCustomerMembershipsRequest models a request for the
// 'customer-list-customer-memberships' operation. See [1] for more information.
//
// List CustomerMemberships belonging to the executing user.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-list-customer-memberships
type ListCustomerMembershipsRequest struct {
	Limit *int64
	Skip  *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListCustomerMembershipsRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListCustomerMembershipsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListCustomerMembershipsRequest) url() string {
	u := url.URL{
		Path:     "/v2/customer-memberships",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListCustomerMembershipsRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
