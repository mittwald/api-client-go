package backupclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteProjectBackupExportRequest models a request for the
// 'backup-delete-project-backup-export' operation. See [1] for more information.
//
// Delete a ProjectBackupExport.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-delete-project-backup-export
type DeleteProjectBackupExportRequest struct {
	ProjectBackupID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteProjectBackupExportRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeleteProjectBackupExportRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteProjectBackupExportRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/project-backups/%s/export", url.PathEscape(r.ProjectBackupID)),
	}
	return u.String()
}

func (r *DeleteProjectBackupExportRequest) query() url.Values {
	return nil
}
