package databaseclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListMysqlUsersRequest models a request for the 'database-list-mysql-users'
// operation. See [1] for more information.
//
// List MySQLUsers belonging to a Database.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-list-mysql-users
type ListMysqlUsersRequest struct {
	MysqlDatabaseID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListMysqlUsersRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListMysqlUsersRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListMysqlUsersRequest) url() string {
	return fmt.Sprintf("/v2/mysql-databases/%s/users", url.PathEscape(r.MysqlDatabaseID))
}

func (r *ListMysqlUsersRequest) query() url.Values {
	return nil
}
