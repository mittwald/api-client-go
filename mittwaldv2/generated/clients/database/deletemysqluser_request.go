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

type DeleteMysqlUserRequest struct {
	MysqlUserID uuid.UUID
}

func (r *DeleteMysqlUserRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteMysqlUserRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteMysqlUserRequest) url() string {
	return fmt.Sprintf("/v2/mysql-users/%s", url.PathEscape(r.MysqlUserID.String()))
}

func (r *DeleteMysqlUserRequest) query() url.Values {
	return nil
}