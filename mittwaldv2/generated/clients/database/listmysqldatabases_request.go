package database

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListMysqlDatabasesRequest models a request for the
// 'database-list-mysql-databases' operation. See [1] for more information.
//
// List MySQLDatabases belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-list-mysql-databases
type ListMysqlDatabasesRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListMysqlDatabasesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListMysqlDatabasesRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListMysqlDatabasesRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/mysql-databases", url.PathEscape(r.ProjectID))
}

func (r *ListMysqlDatabasesRequest) query() url.Values {
	return nil
}
