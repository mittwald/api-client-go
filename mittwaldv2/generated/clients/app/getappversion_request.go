package app

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
func (r *GetAppversionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetAppversionRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetAppversionRequest) url() string {
	return fmt.Sprintf("/v2/apps/%s/versions/%s", url.PathEscape(r.AppID), url.PathEscape(r.AppVersionID))
}

func (r *GetAppversionRequest) query() url.Values {
	return nil
}
