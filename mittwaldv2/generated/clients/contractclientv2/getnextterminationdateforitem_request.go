package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetNextTerminationDateForItemRequest models a request for the
// 'contract-get-next-termination-date-for-item' operation. See [1] for more
// information.
//
// Return the next TerminationDate for the ContractItem with the given ID.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/contract-get-next-termination-date-for-item
type GetNextTerminationDateForItemRequest struct {
	ContractID     string
	ContractItemID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetNextTerminationDateForItemRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetNextTerminationDateForItemRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetNextTerminationDateForItemRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contracts/%s/items/%s/next-termination-dates", url.PathEscape(r.ContractID), url.PathEscape(r.ContractItemID)),
	}
	return u.String()
}

func (r *GetNextTerminationDateForItemRequest) query() url.Values {
	return nil
}
