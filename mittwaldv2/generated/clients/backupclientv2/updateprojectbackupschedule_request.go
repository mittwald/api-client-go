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

// UpdateProjectBackupScheduleRequest models a request for the
// 'backup-update-project-backup-schedule' operation. See [1] for more information.
//
// Update a ProjectBackupSchedule.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/backup/backup-update-project-backup-schedule
type UpdateProjectBackupScheduleRequest struct {
	Body                    UpdateProjectBackupScheduleRequestBody
	ProjectBackupScheduleID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateProjectBackupScheduleRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *UpdateProjectBackupScheduleRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateProjectBackupScheduleRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/project-backup-schedules/%s", url.PathEscape(r.ProjectBackupScheduleID)),
	}
	return u.String()
}

func (r *UpdateProjectBackupScheduleRequest) query() url.Values {
	return nil
}
