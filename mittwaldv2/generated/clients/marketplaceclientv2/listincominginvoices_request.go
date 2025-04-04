package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListIncomingInvoicesRequest models a request for the
// 'contributor-list-incoming-invoices' operation. See [1] for more information.
//
// List incoming Invoices of a Contributor.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/contributor-list-incoming-invoices
type ListIncomingInvoicesRequest struct {
	ContributorID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListIncomingInvoicesRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListIncomingInvoicesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListIncomingInvoicesRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contributors/%s/invoices/incoming", url.PathEscape(r.ContributorID)),
	}
	return u.String()
}

func (r *ListIncomingInvoicesRequest) query() url.Values {
	return nil
}
