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

// UpdateCustomerRequest models a request for the 'customer-update-customer'
// operation. See [1] for more information.
//
// Update a customer profile.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-update-customer
type UpdateCustomerRequest struct {
	Body       UpdateCustomerRequestBody
	CustomerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateCustomerRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *UpdateCustomerRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateCustomerRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s", url.PathEscape(r.CustomerID))
}

func (r *UpdateCustomerRequest) query() url.Values {
	return nil
}
