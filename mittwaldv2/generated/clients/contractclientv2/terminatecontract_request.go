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

// TerminateContractRequest models a request for the 'contract-terminate-contract'
// operation. See [1] for more information.
//
// Schedule the Termination of a Contract.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-terminate-contract
type TerminateContractRequest struct {
	Body       TerminateContractRequestBody
	ContractID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *TerminateContractRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *TerminateContractRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *TerminateContractRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contracts/%s/termination", url.PathEscape(r.ContractID)),
	}
	return u.String()
}

func (r *TerminateContractRequest) query() url.Values {
	return nil
}
