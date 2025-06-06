package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetSystemsoftwareversionRequest models a request for the
// 'app-get-systemsoftwareversion' operation. See [1] for more information.
//
// Get a SystemSoftwareVersion.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/app/app-get-systemsoftwareversion
type GetSystemsoftwareversionRequest struct {
	SystemSoftwareID        string
	SystemSoftwareVersionID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetSystemsoftwareversionRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetSystemsoftwareversionRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetSystemsoftwareversionRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/system-softwares/%s/versions/%s", url.PathEscape(r.SystemSoftwareID), url.PathEscape(r.SystemSoftwareVersionID)),
	}
	return u.String()
}

func (r *GetSystemsoftwareversionRequest) query() url.Values {
	return nil
}
