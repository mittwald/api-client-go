package domainclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ResendDomainEmailRequest models a request for the 'domain-resend-domain-email'
// operation. See [1] for more information.
//
// Resend a Domain email.
//
// Trigger a resend of a confirmation or registrant verification email. Has no
// effect on .de Domains.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/domain-resend-domain-email
type ResendDomainEmailRequest struct {
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ResendDomainEmailRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ResendDomainEmailRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ResendDomainEmailRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/domains/%s/actions/resend-email", url.PathEscape(r.DomainID)),
	}
	return u.String()
}

func (r *ResendDomainEmailRequest) query() url.Values {
	return nil
}
