package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListProjectInvitesRequest models a request for the
// 'project-list-project-invites' operation. See [1] for more information.
//
// List ProjectInvites belonging to the executing user.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-list-project-invites
type ListProjectInvitesRequest struct {
	Limit *int64
	Skip  *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListProjectInvitesRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListProjectInvitesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListProjectInvitesRequest) url() string {
	u := url.URL{
		Path:     "/v2/project-invites",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListProjectInvitesRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	return q
}
