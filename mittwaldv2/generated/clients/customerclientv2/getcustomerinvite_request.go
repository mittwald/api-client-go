package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetCustomerInviteRequest models a request for the 'customer-get-customer-invite'
// operation. See [1] for more information.
//
// Get a CustomerInvite.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-get-customer-invite
type GetCustomerInviteRequest struct {
	CustomerInviteID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetCustomerInviteRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
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

func (r *GetCustomerInviteRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetCustomerInviteRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customer-invites/%s", url.PathEscape(r.CustomerInviteID)),
	}
	return u.String()
}

func (r *GetCustomerInviteRequest) query() url.Values {
	return nil
}
