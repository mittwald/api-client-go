package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListInvitesForProjectRequest models a request for the
// 'project-list-invites-for-project' operation. See [1] for more information.
//
// List Invites belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-list-invites-for-project
type ListInvitesForProjectRequest struct {
	ProjectID string
	Limit     *int64
	Skip      *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListInvitesForProjectRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *ListInvitesForProjectRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListInvitesForProjectRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/invites", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListInvitesForProjectRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
