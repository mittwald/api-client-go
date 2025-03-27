package domainclientv2

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

// DeprecatedChangeOwnercOfDomainRequest models a request for the
// 'deprecated-domain-change-ownerc-of-domain' operation. See [1] for more
// information.
//
// Change the owner contact of a domain.
//
// This operation is deprecated. Use the PATCH
// /v2/domains/{domainId}/contacts/{contact} endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/deprecated-domain-change-ownerc-of-domain
type DeprecatedChangeOwnercOfDomainRequest struct {
	Body     DeprecatedChangeOwnercOfDomainRequestBody
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedChangeOwnercOfDomainRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *DeprecatedChangeOwnercOfDomainRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedChangeOwnercOfDomainRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/domains/%s/handles/ownerc", url.PathEscape(r.DomainID)),
	}
	return u.String()
}

func (r *DeprecatedChangeOwnercOfDomainRequest) query() url.Values {
	return nil
}
