package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetInstalledSystemsoftwareForAppinstallationRequest models a request for the
// 'app-get-installed-systemsoftware-for-appinstallation' operation. See [1] for
// more information.
//
// Get the installed `SystemSoftware' for a specific `AppInstallation`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/app/app-get-installed-systemsoftware-for-appinstallation
type GetInstalledSystemsoftwareForAppinstallationRequest struct {
	AppInstallationID string
	TagFilter         *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetInstalledSystemsoftwareForAppinstallationRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetInstalledSystemsoftwareForAppinstallationRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetInstalledSystemsoftwareForAppinstallationRequest) url() string {
	return fmt.Sprintf("/v2/app-installations/%s/systemSoftware", url.PathEscape(r.AppInstallationID))
}

func (r *GetInstalledSystemsoftwareForAppinstallationRequest) query() url.Values {
	q := make(url.Values)
	if r.TagFilter != nil {
		q.Set("tagFilter", *r.TagFilter)
	}
	return q
}
