package contract

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListContractsRequest models a request for the 'contract-list-contracts'
// operation. See [1] for more information.
//
// Return a list of Contracts for the given Customer.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-list-contracts
type ListContractsRequest struct {
	CustomerID string
	Limit      *int64
	Skip       *int64
	Page       *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListContractsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListContractsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListContractsRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/contracts", url.PathEscape(r.CustomerID))
}

func (r *ListContractsRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	return q
}
