package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetAppversionRequest models a request for the 'app-get-appversion' operation.
// See [1] for more information.
//
// Get an AppVersion.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/app/app-get-appversion
type GetAppversionRequest struct {
	AppID        string
	AppVersionID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetAppversionRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetAppversionRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetAppversionRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/apps/%s/versions/%s", url.PathEscape(r.AppID), url.PathEscape(r.AppVersionID)),
	}
	return u.String()
}

func (r *GetAppversionRequest) query() url.Values {
	return nil
}
