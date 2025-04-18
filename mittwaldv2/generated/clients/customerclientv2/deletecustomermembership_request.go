package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteCustomerMembershipRequest models a request for the
// 'customer-delete-customer-membership' operation. See [1] for more information.
//
// Delete a CustomerMembership.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-delete-customer-membership
type DeleteCustomerMembershipRequest struct {
	CustomerMembershipID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteCustomerMembershipRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *DeleteCustomerMembershipRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteCustomerMembershipRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customer-memberships/%s", url.PathEscape(r.CustomerMembershipID)),
	}
	return u.String()
}

func (r *DeleteCustomerMembershipRequest) query() url.Values {
	return nil
}
