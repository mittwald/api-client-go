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

type DeleteProjectAvatarRequest struct {
	ProjectID uuid.UUID
}

func (r *DeleteProjectAvatarRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteProjectAvatarRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteProjectAvatarRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/avatar", url.PathEscape(r.ProjectID.String()))
}

func (r *DeleteProjectAvatarRequest) query() url.Values {
	return nil
}
