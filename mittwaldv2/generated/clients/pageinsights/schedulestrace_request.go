package pageinsights

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

// ScheduleStraceRequest models a request for the 'pageinsights-schedule-strace'
// operation. See [1] for more information.
//
// Schedule a strace measurement for a single http request.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/page
// insights/pageinsights-schedule-strace
type ScheduleStraceRequest struct {
	Body      ScheduleStraceRequestBody
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
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
	return fmt.Sprintf("/v2/projects/%s/straces", url.PathEscape(r.ProjectID))
}

func (r *ScheduleStraceRequest) query() url.Values {
	return nil
}
