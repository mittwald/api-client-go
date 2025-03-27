package sshsftpuserclientv2

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

// CreateSFTPUserRequest models a request for the 'sftp-user-create-sftp-user'
// operation. See [1] for more information.
//
// Create an SFTPUser for a Project.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/ssh/sftp
// user/sftp-user-create-sftp-user
type CreateSFTPUserRequest struct {
	Body      CreateSFTPUserRequestBody
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateSFTPUserRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
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

func (r *CreateSFTPUserRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateSFTPUserRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/sftp-users", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *CreateSFTPUserRequest) query() url.Values {
	return nil
}
