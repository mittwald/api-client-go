package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetProjectInviteRequest models a request for the 'project-get-project-invite'
// operation. See [1] for more information.
//
// Get a ProjectInvite.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-get-project-invite
type GetProjectInviteRequest struct {
	ProjectInviteID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetProjectInviteRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetProjectInviteRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetProjectInviteRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/project-invites/%s", url.PathEscape(r.ProjectInviteID)),
	}
	return u.String()
}

func (r *GetProjectInviteRequest) query() url.Values {
	return nil
}
