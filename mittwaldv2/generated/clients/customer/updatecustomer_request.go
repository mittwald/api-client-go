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

type UpdateCustomerRequest struct {
	Body       UpdateCustomerRequestBody
	CustomerID string
}

func (r *UpdateCustomerRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *UpdateCustomerRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateCustomerRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s", url.PathEscape(r.CustomerID))
}

func (r *UpdateCustomerRequest) query() url.Values {
	return nil
}