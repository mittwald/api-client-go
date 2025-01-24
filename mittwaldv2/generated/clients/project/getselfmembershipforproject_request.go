package project

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type GetSelfMembershipForProjectRequest struct {
	ProjectID string
}

func (r *GetSelfMembershipForProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetSelfMembershipForProjectRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetSelfMembershipForProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/memberships/self", url.PathEscape(r.ProjectID))
}

func (r *GetSelfMembershipForProjectRequest) query() url.Values {
	return nil
}
