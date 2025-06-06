package databaseclientv2

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
func (r *UpdateRedisDatabaseDescriptionRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
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

func (r *UpdateRedisDatabaseDescriptionRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateRedisDatabaseDescriptionRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/redis-databases/%s/description", url.PathEscape(r.RedisDatabaseID)),
	}
	return u.String()
}

func (r *UpdateRedisDatabaseDescriptionRequest) query() url.Values {
	return nil
}
