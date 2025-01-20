package sshsftpuser

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

type UpdateSFTPUserRequest struct {
	Body       UpdateSFTPUserRequestBody
	SFTPUserID string
}

func (r *UpdateSFTPUserRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateSFTPUserRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateSFTPUserRequest) url() string {
	return fmt.Sprintf("/v2/sftp-users/%s", url.PathEscape(r.SFTPUserID))
}

func (r *UpdateSFTPUserRequest) query() url.Values {
	return nil
}
