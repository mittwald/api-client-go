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

type CreateProjectBackupExportRequest struct {
	Body            CreateProjectBackupExportRequestBody
	ProjectBackupID uuid.UUID
}

func (r *CreateProjectBackupExportRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *CreateProjectBackupExportRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *CreateProjectBackupExportRequest) url() string {
	return fmt.Sprintf("/v2/project-backups/%s/export", url.PathEscape(r.ProjectBackupID.String()))
}

func (r *CreateProjectBackupExportRequest) query() url.Values {
	return nil
}
