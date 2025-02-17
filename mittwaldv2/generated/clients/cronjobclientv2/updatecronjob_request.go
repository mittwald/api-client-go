package cronjobclientv2

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

// UpdateCronjobRequest models a request for the 'cronjob-update-cronjob'
// operation. See [1] for more information.
//
// Update a Cronjob.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-update-cronjob
type UpdateCronjobRequest struct {
	Body      UpdateCronjobRequestBody
	CronjobID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateCronjobRequest) BuildRequest() (*http.Request, error) {
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

func (r *UpdateCronjobRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateCronjobRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/cronjobs/%s", url.PathEscape(r.CronjobID)),
	}
	return u.String()
}

func (r *UpdateCronjobRequest) query() url.Values {
	return nil
}
