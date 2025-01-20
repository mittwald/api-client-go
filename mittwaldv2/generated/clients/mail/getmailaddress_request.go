package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetMailAddressRequest struct {
	MailAddressID uuid.UUID
}

func (r *GetMailAddressRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetMailAddressRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetMailAddressRequest) url() string {
	return fmt.Sprintf("/v2/mail-addresses/%s", url.PathEscape(r.MailAddressID.String()))
}

func (r *GetMailAddressRequest) query() url.Values {
	return nil
}
