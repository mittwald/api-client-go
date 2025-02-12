package mailclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListProjectMailSettingsRequest models a request for the
// 'mail-list-project-mail-settings' operation. See [1] for more information.
//
// List mail settings of a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/mail-list-project-mail-settings
type ListProjectMailSettingsRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListProjectMailSettingsRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListProjectMailSettingsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListProjectMailSettingsRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/mail-settings", url.PathEscape(r.ProjectID))
}

func (r *ListProjectMailSettingsRequest) query() url.Values {
	return nil
}
