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

type UpdateCustomerMembershipRequest struct {
	Body                 UpdateCustomerMembershipRequestBody
	CustomerMembershipID uuid.UUID
}

func (r *UpdateCustomerMembershipRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateCustomerMembershipRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateCustomerMembershipRequest) url() string {
	return fmt.Sprintf("/v2/customer-memberships/%s", url.PathEscape(r.CustomerMembershipID.String()))
}

func (r *UpdateCustomerMembershipRequest) query() url.Values {
	return nil
}
