package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type DeleteIngressRequest struct {
	IngressID uuid.UUID
}

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
	return fmt.Sprintf("/v2/ingresses/%s", url.PathEscape(r.IngressID.String()))
}

func (r *DeleteIngressRequest) query() url.Values {
	return nil
}
