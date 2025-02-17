package contractclientv2

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

// TerminateContractItemRequest models a request for the
// 'contract-terminate-contract-item' operation. See [1] for more information.
//
// Schedule the Termination of a ContractItem.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-terminate-contract-item
type TerminateContractItemRequest struct {
	Body           TerminateContractItemRequestBody
	ContractID     string
	ContractItemID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *TerminateContractItemRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *TerminateContractItemRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *TerminateContractItemRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contracts/%s/items/%s/termination", url.PathEscape(r.ContractID), url.PathEscape(r.ContractItemID)),
	}
	return u.String()
}

func (r *TerminateContractItemRequest) query() url.Values {
	return nil
}
