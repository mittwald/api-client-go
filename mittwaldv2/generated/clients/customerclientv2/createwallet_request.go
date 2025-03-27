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
func (r *CreateWalletRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *CreateWalletRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateWalletRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customers/%s/wallet", url.PathEscape(r.CustomerID)),
	}
	return u.String()
}

func (r *CreateWalletRequest) query() url.Values {
	return nil
}
