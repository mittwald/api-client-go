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

type GetRedisDatabaseRequest struct {
	RedisDatabaseID uuid.UUID
}

func (r *GetRedisDatabaseRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetRedisDatabaseRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetRedisDatabaseRequest) url() string {
	return fmt.Sprintf("/v2/redis-databases/%s", url.PathEscape(r.RedisDatabaseID.String()))
}

func (r *GetRedisDatabaseRequest) query() url.Values {
	return nil
}
