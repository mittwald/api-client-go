package projectclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetProjectTokenInviteRequest models a request for the
// 'project-get-project-token-invite' operation. See [1] for more information.
//
// Get a ProjectInvite by token.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-get-project-token-invite
type GetProjectTokenInviteRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetProjectTokenInviteRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetProjectTokenInviteRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetProjectTokenInviteRequest) url() string {
	u := url.URL{
		Path: "/v2/project-token-invite",
	}
	return u.String()
}

func (r *GetProjectTokenInviteRequest) query() url.Values {
	return nil
}
