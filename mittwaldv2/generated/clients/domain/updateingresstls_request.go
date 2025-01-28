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

// UpdateIngressTLSRequest models a request for the 'ingress-update-ingress-tls'
// operation. See [1] for more information.
//
// Update the tls settings of an Ingress.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/ingress-update-ingress-tls
type UpdateIngressTLSRequest struct {
	Body      UpdateIngressTLSRequestBody
	IngressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateIngressTLSRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateIngressTLSRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateIngressTLSRequest) url() string {
	return fmt.Sprintf("/v2/ingresses/%s/tls", url.PathEscape(r.IngressID))
}

func (r *UpdateIngressTLSRequest) query() url.Values {
	return nil
}
