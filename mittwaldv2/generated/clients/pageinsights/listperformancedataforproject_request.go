package pageinsights

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type ListPerformanceDataForProjectRequest struct {
	ProjectID uuid.UUID
	Domain    *string
}

func (r *ListPerformanceDataForProjectRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListPerformanceDataForProjectRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListPerformanceDataForProjectRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/page-insights", url.PathEscape(r.ProjectID.String()))
}

func (r *ListPerformanceDataForProjectRequest) query() url.Values {
	q := make(url.Values)
	if r.Domain != nil {
		q.Set("domain", *r.Domain)
	}
	return q
}
