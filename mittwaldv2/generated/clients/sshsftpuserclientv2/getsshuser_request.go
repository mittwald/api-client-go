package sshsftpuserclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetSSHUserRequest models a request for the 'ssh-user-get-ssh-user' operation.
// See [1] for more information.
//
// Get an SSHUser.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/ssh/sftp
// user/ssh-user-get-ssh-user
type GetSSHUserRequest struct {
	SSHUserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetSSHUserRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetSSHUserRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetSSHUserRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/ssh-users/%s", url.PathEscape(r.SSHUserID)),
	}
	return u.String()
}

func (r *GetSSHUserRequest) query() url.Values {
	return nil
}
