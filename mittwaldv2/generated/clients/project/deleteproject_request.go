package project

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeleteProjectRequest struct {
	ProjectID string
}

func (r *DeleteProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteProjectRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s", url.PathEscape(r.ProjectID))
}

func (r *DeleteProjectRequest) query() url.Values {
	return nil
}
