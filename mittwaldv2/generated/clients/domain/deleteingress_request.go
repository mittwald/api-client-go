package domain

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
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteIngressRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteIngressRequest) url() string {
	return fmt.Sprintf("/v2/ingresses/%s", url.PathEscape(r.IngressID))
}

func (r *DeleteIngressRequest) query() url.Values {
	return nil
}
