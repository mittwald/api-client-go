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

// AcceptCustomerInviteRequest models a request for the
// 'customer-accept-customer-invite' operation. See [1] for more information.
//
// Accept a CustomerInvite.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-accept-customer-invite
type AcceptCustomerInviteRequest struct {
	Body             AcceptCustomerInviteRequestBody
	CustomerInviteID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AcceptCustomerInviteRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *AcceptCustomerInviteRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *AcceptCustomerInviteRequest) url() string {
	return fmt.Sprintf("/v2/customer-invites/%s/actions/accept", url.PathEscape(r.CustomerInviteID))
}

func (r *AcceptCustomerInviteRequest) query() url.Values {
	return nil
}
