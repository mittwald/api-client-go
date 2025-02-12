package userclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedEditSSHKeyRequest models a request for the
// 'deprecated-user-edit-ssh-key' operation. See [1] for more information.
//
// Edit a stored ssh-key. Replaced by `PUT` `/v2/users/self/ssh-keys/{sshKeyId}`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-edit-ssh-key
type DeprecatedEditSSHKeyRequest struct {
	Body     DeprecatedEditSSHKeyRequestBody
	SSHKeyID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedEditSSHKeyRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedEditSSHKeyRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedEditSSHKeyRequest) url() string {
	return fmt.Sprintf("/v2/signup/ssh/%s", url.PathEscape(r.SSHKeyID))
}

func (r *DeprecatedEditSSHKeyRequest) query() url.Values {
	return nil
}
