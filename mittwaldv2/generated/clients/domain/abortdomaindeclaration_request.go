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

type AbortDomainDeclarationRequest struct {
	DomainID uuid.UUID
}

func (r *AbortDomainDeclarationRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *AbortDomainDeclarationRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *AbortDomainDeclarationRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/declaration", url.PathEscape(r.DomainID.String()))
}

func (r *AbortDomainDeclarationRequest) query() url.Values {
	return nil
}
