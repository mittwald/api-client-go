package project

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetProjectRequest struct {
	ProjectID uuid.UUID
}

func (r *GetProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetProjectRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s", url.PathEscape(r.ProjectID.String()))
}

func (r *GetProjectRequest) query() url.Values {
	return nil
}
