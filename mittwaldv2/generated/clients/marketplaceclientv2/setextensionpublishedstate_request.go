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

// SetExtensionPublishedStateRequest models a request for the
// 'extension-set-extension-published-state' operation. See [1] for more
// information.
//
// Publish or withdraw an Extension.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-set-extension-published-state
type SetExtensionPublishedStateRequest struct {
	Body          SetExtensionPublishedStateRequestBody
	ContributorID string
	ExtensionID   string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *SetExtensionPublishedStateRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *SetExtensionPublishedStateRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *SetExtensionPublishedStateRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contributors/%s/extensions/%s/published", url.PathEscape(r.ContributorID), url.PathEscape(r.ExtensionID)),
	}
	return u.String()
}

func (r *SetExtensionPublishedStateRequest) query() url.Values {
	return nil
}
