package customerclientv2

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

// CreateCustomerRequest models a request for the 'customer-create-customer'
// operation. See [1] for more information.
//
// Create a new customer profile.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-create-customer
type CreateCustomerRequest struct {
	Body CreateCustomerRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateCustomerRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *CreateCustomerRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *CreateCustomerRequest) url() string {
	return "/v2/customers"
}

func (r *CreateCustomerRequest) query() url.Values {
	return nil
}
