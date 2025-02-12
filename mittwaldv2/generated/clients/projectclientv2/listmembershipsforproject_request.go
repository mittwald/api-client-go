package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListMembershipsForProjectRequest models a request for the
// 'project-list-memberships-for-project' operation. See [1] for more information.
//
// List Memberships belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-list-memberships-for-project
type ListMembershipsForProjectRequest struct {
	ProjectID string
	Limit     *int64
	Skip      *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListMembershipsForProjectRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *ListMembershipsForProjectRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListMembershipsForProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/memberships", url.PathEscape(r.ProjectID))
}

func (r *ListMembershipsForProjectRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
