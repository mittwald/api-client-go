package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetDetailOfContractByCertificateRequest models a request for the
// 'contract-get-detail-of-contract-by-certificate' operation. See [1] for more
// information.
//
// Return the Contract for the given Certificate.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-get-detail-of-contract-by-certificate
type GetDetailOfContractByCertificateRequest struct {
	CertificateID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetDetailOfContractByCertificateRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetDetailOfContractByCertificateRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetDetailOfContractByCertificateRequest) url() string {
	return fmt.Sprintf("/v2/certificates/%s/contract", url.PathEscape(r.CertificateID))
}

func (r *GetDetailOfContractByCertificateRequest) query() url.Values {
	return nil
}
