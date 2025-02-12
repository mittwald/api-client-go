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

// AbortExecutionRequest models a request for the 'cronjob-abort-execution'
// operation. See [1] for more information.
//
// Abort a CronjobExecution.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-abort-execution
type AbortExecutionRequest struct {
	Body        any
	CronjobID   string
	ExecutionID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AbortExecutionRequest) BuildRequest() (*http.Request, error) {
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

func (r *AbortExecutionRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *AbortExecutionRequest) url() string {
	return fmt.Sprintf("/v2/cronjobs/%s/executions/%s/actions/abort", url.PathEscape(r.CronjobID), url.PathEscape(r.ExecutionID))
}

func (r *AbortExecutionRequest) query() url.Values {
	return nil
}
