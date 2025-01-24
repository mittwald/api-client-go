package app

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetAppRequest struct {
	AppID uuid.UUID
}

func (r *GetAppRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetAppRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetAppRequest) url() string {
	return fmt.Sprintf("/v2/apps/%s", url.PathEscape(r.AppID.String()))
}

func (r *GetAppRequest) query() url.Values {
	return nil
}
