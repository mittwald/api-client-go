package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedDeleteSSHKeyRequest models a request for the
// 'deprecated-user-delete-ssh-key' operation. See [1] for more information.
//
// Remove a ssh-key. Replaced by `DELETE` `/v2/users/self/ssh-keys/{sshKeyId}`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-delete-ssh-key
type DeprecatedDeleteSSHKeyRequest struct {
	SSHKeyID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedDeleteSSHKeyRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedDeleteSSHKeyRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedDeleteSSHKeyRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/signup/ssh/%s", url.PathEscape(r.SSHKeyID)),
	}
	return u.String()
}

func (r *DeprecatedDeleteSSHKeyRequest) query() url.Values {
	return nil
}
