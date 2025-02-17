package domainclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListCertificatesRequest models a request for the 'ssl-list-certificates'
// operation. See [1] for more information.
//
// List Certificates belonging to a Project or an Ingress.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/ssl-list-certificates
type ListCertificatesRequest struct {
	ProjectID *string
	IngressID *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListCertificatesRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListCertificatesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListCertificatesRequest) url() string {
	u := url.URL{
		Path:     "/v2/certificates",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListCertificatesRequest) query() url.Values {
	q := make(url.Values)
	if r.ProjectID != nil {
		q.Set("projectId", *r.ProjectID)
	}
	if r.IngressID != nil {
		q.Set("ingressId", *r.IngressID)
	}
	return q
}
