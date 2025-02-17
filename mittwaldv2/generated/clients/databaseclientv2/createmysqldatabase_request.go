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

// CreateMysqlDatabaseRequest models a request for the
// 'database-create-mysql-database' operation. See [1] for more information.
//
// Create a MySQLDatabase with a MySQLUser.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-create-mysql-database
type CreateMysqlDatabaseRequest struct {
	Body      CreateMysqlDatabaseRequestBody
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateMysqlDatabaseRequest) BuildRequest() (*http.Request, error) {
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

func (r *CreateMysqlDatabaseRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateMysqlDatabaseRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/mysql-databases", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *CreateMysqlDatabaseRequest) query() url.Values {
	return nil
}
