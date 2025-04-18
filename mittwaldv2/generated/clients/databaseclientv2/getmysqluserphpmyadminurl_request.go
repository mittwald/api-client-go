package databaseclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetMysqlUserPhpMyAdminURLRequest models a request for the
// 'database-get-mysql-user-php-my-admin-url' operation. See [1] for more
// information.
//
// Get a MySQLUser's PhpMyAdmin-URL.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-get-mysql-user-php-my-admin-url
type GetMysqlUserPhpMyAdminURLRequest struct {
	MysqlUserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetMysqlUserPhpMyAdminURLRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *GetMysqlUserPhpMyAdminURLRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetMysqlUserPhpMyAdminURLRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/mysql-users/%s/php-my-admin-url", url.PathEscape(r.MysqlUserID)),
	}
	return u.String()
}

func (r *GetMysqlUserPhpMyAdminURLRequest) query() url.Values {
	return nil
}
