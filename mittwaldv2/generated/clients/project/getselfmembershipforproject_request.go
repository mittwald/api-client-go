package project

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetSelfMembershipForProjectRequest models a request for the
// 'project-get-self-membership-for-project' operation. See [1] for more
// information.
//
// Get the executing user's membership in a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-get-self-membership-for-project
type GetSelfMembershipForProjectRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
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
