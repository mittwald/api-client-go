package customerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// IsCustomerLegallyCompetentRequest models a request for the
// 'customer-is-customer-legally-competent' operation. See [1] for more
// information.
//
// Check if the customer profile has a valid contract partner configured.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/customer/customer-is-customer-legally-competent
type IsCustomerLegallyCompetentRequest struct {
	CustomerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *IsCustomerLegallyCompetentRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *IsCustomerLegallyCompetentRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *IsCustomerLegallyCompetentRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customers/%s/legally-competent", url.PathEscape(r.CustomerID)),
	}
	return u.String()
}

func (r *IsCustomerLegallyCompetentRequest) query() url.Values {
	return nil
}
