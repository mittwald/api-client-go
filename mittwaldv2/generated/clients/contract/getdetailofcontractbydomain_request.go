package contract

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
func (r *GetDetailOfContractByDomainRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetDetailOfContractByDomainRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetDetailOfContractByDomainRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/contract", url.PathEscape(r.DomainID))
}

func (r *GetDetailOfContractByDomainRequest) query() url.Values {
	return nil
}
