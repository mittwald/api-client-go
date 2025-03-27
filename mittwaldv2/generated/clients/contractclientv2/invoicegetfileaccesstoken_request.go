package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// InvoiceGetFileAccessTokenRequest models a request for the
// 'invoice-get-file-access-token' operation. See [1] for more information.
//
// Request an Access Token for the Invoice file.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/invoice-get-file-access-token
type InvoiceGetFileAccessTokenRequest struct {
	CustomerID string
	InvoiceID  string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *InvoiceGetFileAccessTokenRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *InvoiceGetFileAccessTokenRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *InvoiceGetFileAccessTokenRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customers/%s/invoices/%s/file-access-token", url.PathEscape(r.CustomerID), url.PathEscape(r.InvoiceID)),
	}
	return u.String()
}

func (r *InvoiceGetFileAccessTokenRequest) query() url.Values {
	return nil
}
