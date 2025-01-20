package project

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type StoragespaceGetServerStatisticsRequest struct {
	ServerID string
}

func (r *StoragespaceGetServerStatisticsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *StoragespaceGetServerStatisticsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *StoragespaceGetServerStatisticsRequest) url() string {
	return fmt.Sprintf("/v2/servers/%s/storage-space-statistics", url.PathEscape(r.ServerID))
}

func (r *StoragespaceGetServerStatisticsRequest) query() url.Values {
	return nil
}
