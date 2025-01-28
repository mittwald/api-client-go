package customer

import (
	"bytes"
	"encoding/json"
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
	Body             any
	CustomerInviteID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ResendCustomerInviteMailRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ResendCustomerInviteMailRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ResendCustomerInviteMailRequest) url() string {
	return fmt.Sprintf("/v2/customer-invites/%s/actions/resend", url.PathEscape(r.CustomerInviteID))
}

func (r *ResendCustomerInviteMailRequest) query() url.Values {
	return nil
}
