package backupclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteProjectBackupScheduleRequest models a request for the
// 'backup-delete-project-backup-schedule' operation. See [1] for more information.
//
// Delete a ProjectBackupSchedule.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-delete-project-backup-schedule
type DeleteProjectBackupScheduleRequest struct {
	ProjectBackupScheduleID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteProjectBackupScheduleRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteProjectBackupScheduleRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteProjectBackupScheduleRequest) url() string {
	return fmt.Sprintf("/v2/project-backup-schedules/%s", url.PathEscape(r.ProjectBackupScheduleID))
}

func (r *DeleteProjectBackupScheduleRequest) query() url.Values {
	return nil
}
