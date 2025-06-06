package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ResendCustomerInviteMailRequest models a request for the
// 'customer-resend-customer-invite-mail' operation. See [1] for more information.
//
// Resend the mail for a CustomerInvite.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-resend-customer-invite-mail
type ResendCustomerInviteMailRequest struct {
	CustomerInviteID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ResendCustomerInviteMailRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
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

func (r *ResendCustomerInviteMailRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ResendCustomerInviteMailRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customer-invites/%s/actions/resend", url.PathEscape(r.CustomerInviteID)),
	}
	return u.String()
}

func (r *ResendCustomerInviteMailRequest) query() url.Values {
	return nil
}
