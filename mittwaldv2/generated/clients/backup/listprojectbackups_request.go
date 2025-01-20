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

type ListProjectBackupsRequest struct {
	ProjectID uuid.UUID
}

func (r *ListProjectBackupsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListProjectBackupsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListProjectBackupsRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/backups", url.PathEscape(r.ProjectID.String()))
}

func (r *ListProjectBackupsRequest) query() url.Values {
	return nil
}