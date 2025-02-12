package domainclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetIngressRequest models a request for the 'ingress-get-ingress' operation. See
// [1] for more information.
//
// Get an Ingress.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/domain/ingress-get-ingress
type GetIngressRequest struct {
	IngressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetIngressRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetIngressRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetIngressRequest) url() string {
	return fmt.Sprintf("/v2/ingresses/%s", url.PathEscape(r.IngressID))
}

func (r *GetIngressRequest) query() url.Values {
	return nil
}
