package backupclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListProjectBackupsRequest models a request for the 'backup-list-project-backups'
// operation. See [1] for more information.
//
// List Backups belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-list-project-backups
type ListProjectBackupsRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListProjectBackupsRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListProjectBackupsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListProjectBackupsRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/backups", url.PathEscape(r.ProjectID))
}

func (r *ListProjectBackupsRequest) query() url.Values {
	return nil
}
