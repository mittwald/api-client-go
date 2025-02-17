package domainclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteIngressRequest models a request for the 'ingress-delete-ingress'
// operation. See [1] for more information.
//
// Delete an Ingress.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/ingress-delete-ingress
type DeleteIngressRequest struct {
	IngressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteIngressRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeleteIngressRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteIngressRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/ingresses/%s", url.PathEscape(r.IngressID)),
	}
	return u.String()
}

func (r *DeleteIngressRequest) query() url.Values {
	return nil
}
