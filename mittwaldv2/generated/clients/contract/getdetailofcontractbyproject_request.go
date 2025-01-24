package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetDetailOfContractByProjectRequest struct {
	ProjectID uuid.UUID
}

func (r *GetDetailOfContractByProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetDetailOfContractByProjectRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetDetailOfContractByProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/contract", url.PathEscape(r.ProjectID.String()))
}

func (r *GetDetailOfContractByProjectRequest) query() url.Values {
	return nil
}
