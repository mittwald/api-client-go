package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// UninstallAppinstallationRequest models a request for the
// 'app-uninstall-appinstallation' operation. See [1] for more information.
//
// Trigger an uninstallation process for an AppInstallation.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/app/app-uninstall-appinstallation
type UninstallAppinstallationRequest struct {
	AppInstallationID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UninstallAppinstallationRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *UninstallAppinstallationRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *UninstallAppinstallationRequest) url() string {
	return fmt.Sprintf("/v2/app-installations/%s", url.PathEscape(r.AppInstallationID))
}

func (r *UninstallAppinstallationRequest) query() url.Values {
	return nil
}
