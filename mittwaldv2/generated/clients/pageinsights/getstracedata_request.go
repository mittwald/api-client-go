package pageinsights

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetStraceDataRequest struct {
	StraceID  uuid.UUID
	ProjectID uuid.UUID
}

func (r *GetStraceDataRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetStraceDataRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetStraceDataRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/straces/%s", url.PathEscape(r.ProjectID.String()), url.PathEscape(r.StraceID.String()))
}

func (r *GetStraceDataRequest) query() url.Values {
	return nil
}