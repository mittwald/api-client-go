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

// CreateExecutionRequest models a request for the 'cronjob-create-execution'
// operation. See [1] for more information.
//
// Trigger a Cronjob.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-create-execution
type CreateExecutionRequest struct {
	Body      any
	CronjobID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateExecutionRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *CreateExecutionRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateExecutionRequest) url() string {
	return fmt.Sprintf("/v2/cronjobs/%s/executions", url.PathEscape(r.CronjobID))
}

func (r *CreateExecutionRequest) query() url.Values {
	return nil
}
