package domainclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListTldContactSchemasRequest models a request for the
// 'domain-list-tld-contact-schemas' operation. See [1] for more information.
//
// List the contact schemas for a TLD.
//
// List the contact schemas describing the fields required to register/transfer a
// Domain.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/domain-list-tld-contact-schemas
type ListTldContactSchemasRequest struct {
	Tld string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListTldContactSchemasRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListTldContactSchemasRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListTldContactSchemasRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/domain-tlds/%s/contact-schemas", url.PathEscape(r.Tld)),
	}
	return u.String()
}

func (r *ListTldContactSchemasRequest) query() url.Values {
	return nil
}
