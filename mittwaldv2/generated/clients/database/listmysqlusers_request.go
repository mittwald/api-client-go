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

type ListMysqlUsersRequest struct {
	MysqlDatabaseID uuid.UUID
}

func (r *ListMysqlUsersRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListMysqlUsersRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListMysqlUsersRequest) url() string {
	return fmt.Sprintf("/v2/mysql-databases/%s/users", url.PathEscape(r.MysqlDatabaseID.String()))
}

func (r *ListMysqlUsersRequest) query() url.Values {
	return nil
}