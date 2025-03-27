package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetExtensionInstanceForCustomerRequest models a request for the
// 'extension-get-extension-instance-for-customer' operation. See [1] for more
// information.
//
// Get the ExtensionInstance of a specific customer and extension, if existing.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-get-extension-instance-for-customer
type GetExtensionInstanceForCustomerRequest struct {
	CustomerID  string
	ExtensionID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetExtensionInstanceForCustomerRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetExtensionInstanceForCustomerRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetExtensionInstanceForCustomerRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/customers/%s/extensions/%s", url.PathEscape(r.CustomerID), url.PathEscape(r.ExtensionID)),
	}
	return u.String()
}

func (r *GetExtensionInstanceForCustomerRequest) query() url.Values {
	return nil
}
