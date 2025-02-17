package cronjobclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/cronjobv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// CreateCronjobRequest models a request for the 'cronjob-create-cronjob'
// operation. See [1] for more information.
//
// Create a Cronjob.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-create-cronjob
type CreateCronjobRequest struct {
	Body      cronjobv2.CronjobRequest
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateCronjobRequest) BuildRequest() (*http.Request, error) {
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

func (r *CreateCronjobRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateCronjobRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/cronjobs", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *CreateCronjobRequest) query() url.Values {
	return nil
}
