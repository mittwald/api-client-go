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

type ResendDomainEmailRequest struct {
	DomainID uuid.UUID
}

func (r *ResendDomainEmailRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ResendDomainEmailRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ResendDomainEmailRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/actions/resend-email", url.PathEscape(r.DomainID.String()))
}

func (r *ResendDomainEmailRequest) query() url.Values {
	return nil
}
