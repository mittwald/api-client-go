package databaseclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteRedisDatabaseRequest models a request for the
// 'database-delete-redis-database' operation. See [1] for more information.
//
// Delete a RedisDatabase.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-delete-redis-database
type DeleteRedisDatabaseRequest struct {
	RedisDatabaseID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteRedisDatabaseRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
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

func (r *DeleteRedisDatabaseRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteRedisDatabaseRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/redis-databases/%s", url.PathEscape(r.RedisDatabaseID)),
	}
	return u.String()
}

func (r *DeleteRedisDatabaseRequest) query() url.Values {
	return nil
}
