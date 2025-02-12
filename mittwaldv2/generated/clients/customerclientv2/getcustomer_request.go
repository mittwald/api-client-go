package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetCustomerRequest models a request for the 'customer-get-customer' operation.
// See [1] for more information.
//
// Get a customer profile.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-get-customer
type GetCustomerRequest struct {
	CustomerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetCustomerRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetCustomerRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetCustomerRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s", url.PathEscape(r.CustomerID))
}

func (r *GetCustomerRequest) query() url.Values {
	return nil
}
