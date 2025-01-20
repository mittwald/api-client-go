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

type GetMysqlDatabaseRequest struct {
	MysqlDatabaseID uuid.UUID
}

func (r *GetMysqlDatabaseRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetMysqlDatabaseRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetMysqlDatabaseRequest) url() string {
	return fmt.Sprintf("/v2/mysql-databases/%s", url.PathEscape(r.MysqlDatabaseID.String()))
}

func (r *GetMysqlDatabaseRequest) query() url.Values {
	return nil
}
