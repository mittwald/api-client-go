package pageinsights

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

type ScheduleStraceRequest struct {
	Body      ScheduleStraceRequestBody
	ProjectID uuid.UUID
}

func (r *ScheduleStraceRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ScheduleStraceRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ScheduleStraceRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/straces", url.PathEscape(r.ProjectID.String()))
}

func (r *ScheduleStraceRequest) query() url.Values {
	return nil
}
