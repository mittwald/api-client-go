package database

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListRedisVersionsRequest models a request for the 'database-list-redis-versions'
// operation. See [1] for more information.
//
// List RedisVersions.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-list-redis-versions
type ListRedisVersionsRequest struct {
	ProjectID *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListRedisVersionsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListRedisVersionsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListRedisVersionsRequest) url() string {
	return "/v2/redis-versions"
}

func (r *ListRedisVersionsRequest) query() url.Values {
	q := make(url.Values)
	if r.ProjectID != nil {
		q.Set("projectId", *r.ProjectID)
	}
	return q
}
