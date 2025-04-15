package marketplaceclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// CancelExtensionTerminationRequest models a request for the
// 'extension-cancel-extension-termination' operation. See [1] for more
// information.
//
// Cancel an Extension Instance Termination.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-cancel-extension-termination
type CancelExtensionTerminationRequest struct {
	Body                CancelExtensionTerminationRequestBody
	ExtensionInstanceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CancelExtensionTerminationRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
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

func (r *CancelExtensionTerminationRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CancelExtensionTerminationRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/extension-instances/%s/termination", url.PathEscape(r.ExtensionInstanceID)),
	}
	return u.String()
}

func (r *CancelExtensionTerminationRequest) query() url.Values {
	return nil
}
