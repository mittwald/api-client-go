package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// UpdateRedisDatabaseDescriptionRequest models a request for the
// 'database-update-redis-database-description' operation. See [1] for more
// information.
//
// Update a RedisDatabase's description.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/database/database-update-redis-database-description
type UpdateRedisDatabaseDescriptionRequest struct {
	Body            UpdateRedisDatabaseDescriptionRequestBody
	RedisDatabaseID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateRedisDatabaseDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateRedisDatabaseDescriptionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateRedisDatabaseDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/redis-databases/%s/description", url.PathEscape(r.RedisDatabaseID))
}

func (r *UpdateRedisDatabaseDescriptionRequest) query() url.Values {
	return nil
}
