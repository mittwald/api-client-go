package databaseclientv2

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

// UpdateMysqlUserRequest models a request for the 'database-update-mysql-user'
// operation. See [1] for more information.
//
// Update a MySQLUser.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-update-mysql-user
type UpdateMysqlUserRequest struct {
	Body        UpdateMysqlUserRequestBody
	MysqlUserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateMysqlUserRequest) BuildRequest() (*http.Request, error) {
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

func (r *UpdateMysqlUserRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateMysqlUserRequest) url() string {
	return fmt.Sprintf("/v2/mysql-users/%s", url.PathEscape(r.MysqlUserID))
}

func (r *UpdateMysqlUserRequest) query() url.Values {
	return nil
}
