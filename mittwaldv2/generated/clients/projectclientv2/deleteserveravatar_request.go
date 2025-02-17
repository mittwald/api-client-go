package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteServerAvatarRequest models a request for the
// 'project-delete-server-avatar' operation. See [1] for more information.
//
// Delete a Server's avatar.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-delete-server-avatar
type DeleteServerAvatarRequest struct {
	ServerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteServerAvatarRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeleteServerAvatarRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteServerAvatarRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/servers/%s/avatar", url.PathEscape(r.ServerID)),
	}
	return u.String()
}

func (r *DeleteServerAvatarRequest) query() url.Values {
	return nil
}
