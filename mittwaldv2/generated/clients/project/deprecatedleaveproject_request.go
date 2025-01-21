package project

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type DeprecatedLeaveProjectRequest struct {
	Body      any
	ProjectID uuid.UUID
}

func (r *DeprecatedLeaveProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *DeprecatedLeaveProjectRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedLeaveProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/leave", url.PathEscape(r.ProjectID.String()))
}

func (r *DeprecatedLeaveProjectRequest) query() url.Values {
	return nil
}