package backupclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// CreateProjectBackupScheduleRequest models a request for the
// 'backup-create-project-backup-schedule' operation. See [1] for more information.
//
// Create a BackupSchedule for a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-create-project-backup-schedule
type CreateProjectBackupScheduleRequest struct {
	Body      CreateProjectBackupScheduleRequestBody
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateProjectBackupScheduleRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *CreateProjectBackupScheduleRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateProjectBackupScheduleRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/backup-schedules", url.PathEscape(r.ProjectID))
}

func (r *CreateProjectBackupScheduleRequest) query() url.Values {
	return nil
}
