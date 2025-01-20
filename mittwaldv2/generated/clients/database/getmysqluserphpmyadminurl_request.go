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

type GetMysqlUserPhpMyAdminURLRequest struct {
	MysqlUserID uuid.UUID
}

func (r *GetMysqlUserPhpMyAdminURLRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetMysqlUserPhpMyAdminURLRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetMysqlUserPhpMyAdminURLRequest) url() string {
	return fmt.Sprintf("/v2/mysql-users/%s/php-my-admin-url", url.PathEscape(r.MysqlUserID.String()))
}

func (r *GetMysqlUserPhpMyAdminURLRequest) query() url.Values {
	return nil
}
