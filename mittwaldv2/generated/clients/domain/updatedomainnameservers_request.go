package domain

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

// UpdateDomainNameserversRequest models a request for the
// 'domain-update-domain-nameservers' operation. See [1] for more information.
//
// Update the nameservers of a Domain.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/domain-update-domain-nameservers
type UpdateDomainNameserversRequest struct {
	Body     UpdateDomainNameserversRequestBody
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateDomainNameserversRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateDomainNameserversRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateDomainNameserversRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/nameservers", url.PathEscape(r.DomainID))
}

func (r *UpdateDomainNameserversRequest) query() url.Values {
	return nil
}
