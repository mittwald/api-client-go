package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type InvoiceGetFileAccessTokenRequest struct {
	CustomerID string
	InvoiceID  uuid.UUID
}

func (r *InvoiceGetFileAccessTokenRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *InvoiceGetFileAccessTokenRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *InvoiceGetFileAccessTokenRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/invoices/%s/file-access-token", url.PathEscape(r.CustomerID), url.PathEscape(r.InvoiceID.String()))
}

func (r *InvoiceGetFileAccessTokenRequest) query() url.Values {
	return nil
}