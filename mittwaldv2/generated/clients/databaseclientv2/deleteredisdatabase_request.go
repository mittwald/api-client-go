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
func (r *DeleteRedisDatabaseRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeleteRedisDatabaseRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteRedisDatabaseRequest) url() string {
	return fmt.Sprintf("/v2/redis-databases/%s", url.PathEscape(r.RedisDatabaseID))
}

func (r *DeleteRedisDatabaseRequest) query() url.Values {
	return nil
}
