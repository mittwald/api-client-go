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

type UpdateProjectBackupDescriptionRequest struct {
	Body            UpdateProjectBackupDescriptionRequestBody
	ProjectBackupID uuid.UUID
}

func (r *UpdateProjectBackupDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateProjectBackupDescriptionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateProjectBackupDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/project-backups/%s/description", url.PathEscape(r.ProjectBackupID.String()))
}

func (r *UpdateProjectBackupDescriptionRequest) query() url.Values {
	return nil
}
