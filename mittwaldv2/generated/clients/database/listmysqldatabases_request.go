package database

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type ListMysqlDatabasesRequest struct {
	ProjectID uuid.UUID
}

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
	return fmt.Sprintf("/v2/projects/%s/mysql-databases", url.PathEscape(r.ProjectID.String()))
}

func (r *ListMysqlDatabasesRequest) query() url.Values {
	return nil
}
