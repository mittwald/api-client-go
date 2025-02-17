package backupclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListProjectBackupSchedulesRequest models a request for the
// 'backup-list-project-backup-schedules' operation. See [1] for more information.
//
// List BackupSchedules belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-list-project-backup-schedules
type ListProjectBackupSchedulesRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListProjectBackupSchedulesRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListProjectBackupSchedulesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListProjectBackupSchedulesRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/backup-schedules", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *ListProjectBackupSchedulesRequest) query() url.Values {
	return nil
}
