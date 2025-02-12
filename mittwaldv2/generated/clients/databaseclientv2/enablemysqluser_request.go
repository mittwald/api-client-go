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

// EnableMysqlUserRequest models a request for the 'database-enable-mysql-user'
// operation. See [1] for more information.
//
// Enable a MySQLUser.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-enable-mysql-user
type EnableMysqlUserRequest struct {
	Body        any
	MysqlUserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *EnableMysqlUserRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *EnableMysqlUserRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *EnableMysqlUserRequest) url() string {
	return fmt.Sprintf("/v2/mysql-users/%s/actions/enable", url.PathEscape(r.MysqlUserID))
}

func (r *EnableMysqlUserRequest) query() url.Values {
	return nil
}
