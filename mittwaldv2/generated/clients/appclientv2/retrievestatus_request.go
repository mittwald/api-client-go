package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// RetrieveStatusRequest models a request for the 'app-retrieve-status' operation.
// See [1] for more information.
//
// Get runtime status belonging to an AppInstallation.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/app/app-retrieve-status
type RetrieveStatusRequest struct {
	AppInstallationID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *RetrieveStatusRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *RetrieveStatusRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *RetrieveStatusRequest) url() string {
	return fmt.Sprintf("/v2/app-installations/%s/status", url.PathEscape(r.AppInstallationID))
}

func (r *RetrieveStatusRequest) query() url.Values {
	return nil
}
