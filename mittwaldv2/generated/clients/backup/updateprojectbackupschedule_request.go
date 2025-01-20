package backup

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type UpdateProjectBackupScheduleRequest struct {
	Body                    UpdateProjectBackupScheduleRequestBody
	ProjectBackupScheduleID uuid.UUID
}

func (r *UpdateProjectBackupScheduleRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateProjectBackupScheduleRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateProjectBackupScheduleRequest) url() string {
	return fmt.Sprintf("/v2/project-backup-schedules/%s", url.PathEscape(r.ProjectBackupScheduleID.String()))
}

func (r *UpdateProjectBackupScheduleRequest) query() url.Values {
	return nil
}