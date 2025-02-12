package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedInvoiceDetailOfInvoiceRequest models a request for the
// 'deprecated-invoice-detail-of-invoice' operation. See [1] for more information.
//
// Get details of an Invoice.
//
// This route is deprecated. Use /v2/invoices/{invoiceId} instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/deprecated-invoice-detail-of-invoice
type DeprecatedInvoiceDetailOfInvoiceRequest struct {
	CustomerID string
	InvoiceID  string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedInvoiceDetailOfInvoiceRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedInvoiceDetailOfInvoiceRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedInvoiceDetailOfInvoiceRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/invoices/%s", url.PathEscape(r.CustomerID), url.PathEscape(r.InvoiceID))
}

func (r *DeprecatedInvoiceDetailOfInvoiceRequest) query() url.Values {
	return nil
}
