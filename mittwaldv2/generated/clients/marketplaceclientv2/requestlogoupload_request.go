package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// RequestLogoUploadRequest models a request for the
// 'extension-request-logo-upload' operation. See [1] for more information.
//
// Add a logo to an extension.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-request-logo-upload
type RequestLogoUploadRequest struct {
	ContributorID string
	ExtensionID   string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *RequestLogoUploadRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
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

func (r *RequestLogoUploadRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *RequestLogoUploadRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contributors/%s/extensions/%s/logo", url.PathEscape(r.ContributorID), url.PathEscape(r.ExtensionID)),
	}
	return u.String()
}

func (r *RequestLogoUploadRequest) query() url.Values {
	return nil
}
