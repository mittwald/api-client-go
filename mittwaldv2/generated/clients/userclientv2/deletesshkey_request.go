package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteSSHKeyRequest models a request for the 'user-delete-ssh-key' operation.
// See [1] for more information.
//
// Remove a ssh-key.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-delete-ssh-key
type DeleteSSHKeyRequest struct {
	SSHKeyID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteSSHKeyRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeleteSSHKeyRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteSSHKeyRequest) url() string {
	return fmt.Sprintf("/v2/users/self/ssh-keys/%s", url.PathEscape(r.SSHKeyID))
}

func (r *DeleteSSHKeyRequest) query() url.Values {
	return nil
}
