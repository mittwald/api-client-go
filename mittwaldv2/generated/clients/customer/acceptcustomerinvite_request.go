package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type AcceptCustomerInviteRequest struct {
	Body             AcceptCustomerInviteRequestBody
	CustomerInviteID uuid.UUID
}

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
	return fmt.Sprintf("/v2/customer-invites/%s/actions/accept", url.PathEscape(r.CustomerInviteID.String()))
}

func (r *AcceptCustomerInviteRequest) query() url.Values {
	return nil
}
