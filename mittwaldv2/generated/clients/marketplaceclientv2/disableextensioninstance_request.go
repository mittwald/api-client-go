package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DisableExtensionInstanceRequest models a request for the
// 'extension-disable-extension-instance' operation. See [1] for more information.
//
// Disable an ExtensionInstance.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-disable-extension-instance
type DisableExtensionInstanceRequest struct {
	ExtensionInstanceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DisableExtensionInstanceRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DisableExtensionInstanceRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DisableExtensionInstanceRequest) url() string {
	return fmt.Sprintf("/v2/extension-instances/%s/actions/disable", url.PathEscape(r.ExtensionInstanceID))
}

func (r *DisableExtensionInstanceRequest) query() url.Values {
	return nil
}
