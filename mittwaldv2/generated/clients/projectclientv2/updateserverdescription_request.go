package projectclientv2

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

// UpdateServerDescriptionRequest models a request for the
// 'project-update-server-description' operation. See [1] for more information.
//
// Update a Servers's description.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-update-server-description
type UpdateServerDescriptionRequest struct {
	Body     UpdateServerDescriptionRequestBody
	ServerID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateServerDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *UpdateServerDescriptionRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateServerDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/servers/%s/description", url.PathEscape(r.ServerID))
}

func (r *UpdateServerDescriptionRequest) query() url.Values {
	return nil
}
