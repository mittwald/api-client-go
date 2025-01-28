package project

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

// StoragespaceReplaceServerNotificationThresholdRequest models a request for the
// 'storagespace-replace-server-notification-threshold' operation. See [1] for more
// information.
//
// Update a Server's storage space notification threshold.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/storagespace-replace-server-notification-threshold
type StoragespaceReplaceServerNotificationThresholdRequest struct {
	Body     StoragespaceReplaceServerNotificationThresholdRequestBody
	ServerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *StoragespaceReplaceServerNotificationThresholdRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *StoragespaceReplaceServerNotificationThresholdRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *StoragespaceReplaceServerNotificationThresholdRequest) url() string {
	return fmt.Sprintf("/v2/servers/%s/storage-space-notification-threshold", url.PathEscape(r.ServerID))
}

func (r *StoragespaceReplaceServerNotificationThresholdRequest) query() url.Values {
	return nil
}
