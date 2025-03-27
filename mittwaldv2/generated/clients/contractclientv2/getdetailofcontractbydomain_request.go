package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetDetailOfContractByDomainRequest models a request for the
// 'contract-get-detail-of-contract-by-domain' operation. See [1] for more
// information.
//
// Return the Contract for the given Domain.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-get-detail-of-contract-by-domain
type GetDetailOfContractByDomainRequest struct {
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetDetailOfContractByDomainRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
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

func (r *GetDetailOfContractByDomainRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetDetailOfContractByDomainRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/domains/%s/contract", url.PathEscape(r.DomainID)),
	}
	return u.String()
}

func (r *GetDetailOfContractByDomainRequest) query() url.Values {
	return nil
}
