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

type GetAppversionRequest struct {
	AppID        uuid.UUID
	AppVersionID uuid.UUID
}

func (r *GetAppversionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetAppversionRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetAppversionRequest) url() string {
	return fmt.Sprintf("/v2/apps/%s/versions/%s", url.PathEscape(r.AppID.String()), url.PathEscape(r.AppVersionID.String()))
}

func (r *GetAppversionRequest) query() url.Values {
	return nil
}