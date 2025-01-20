package customer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type GetCustomerMembershipRequest struct {
	CustomerMembershipID string
}

func (r *GetCustomerMembershipRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetCustomerMembershipRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetCustomerMembershipRequest) url() string {
	return fmt.Sprintf("/v2/customer-memberships/%s", url.PathEscape(r.CustomerMembershipID))
}

func (r *GetCustomerMembershipRequest) query() url.Values {
	return nil
}
