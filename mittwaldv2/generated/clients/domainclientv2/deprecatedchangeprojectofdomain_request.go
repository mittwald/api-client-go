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

// DeprecatedChangeProjectOfDomainRequest models a request for the
// 'deprecated-domain-change-project-of-domain' operation. See [1] for more
// information.
//
// Change the Project relation of a Domain.
//
// This operation is deprecated. Use the PATCH /v2/domains/{domainId}/project-id
// endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/deprecated-domain-change-project-of-domain
type DeprecatedChangeProjectOfDomainRequest struct {
	Body     DeprecatedChangeProjectOfDomainRequestBody
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedChangeProjectOfDomainRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *DeprecatedChangeProjectOfDomainRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedChangeProjectOfDomainRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/domains/%s/projectId", url.PathEscape(r.DomainID)),
	}
	return u.String()
}

func (r *DeprecatedChangeProjectOfDomainRequest) query() url.Values {
	return nil
}
