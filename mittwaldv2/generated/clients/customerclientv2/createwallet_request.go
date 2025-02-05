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

// CreateWalletRequest models a request for the 'customer-create-wallet' operation.
// See [1] for more information.
//
// Create the Wallet for the Customer.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-create-wallet
type CreateWalletRequest struct {
	Body       CreateWalletRequestBody
	CustomerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateWalletRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *CreateWalletRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *CreateWalletRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/wallet", url.PathEscape(r.CustomerID))
}

func (r *CreateWalletRequest) query() url.Values {
	return nil
}
