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

type GetProjectBackupRequest struct {
	ProjectBackupID uuid.UUID
}

func (r *GetProjectBackupRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetProjectBackupRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetProjectBackupRequest) url() string {
	return fmt.Sprintf("/v2/project-backups/%s", url.PathEscape(r.ProjectBackupID.String()))
}

func (r *GetProjectBackupRequest) query() url.Values {
	return nil
}