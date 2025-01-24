package backup

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type DeleteProjectBackupExportRequest struct {
	ProjectBackupID uuid.UUID
}

func (r *DeleteProjectBackupExportRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteProjectBackupExportRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteProjectBackupExportRequest) url() string {
	return fmt.Sprintf("/v2/project-backups/%s/export", url.PathEscape(r.ProjectBackupID.String()))
}

func (r *DeleteProjectBackupExportRequest) query() url.Values {
	return nil
}
