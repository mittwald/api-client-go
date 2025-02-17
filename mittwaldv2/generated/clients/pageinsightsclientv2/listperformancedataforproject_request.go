package pageinsightsclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListPerformanceDataForProjectRequest models a request for the
// 'pageinsights-list-performance-data-for-project' operation. See [1] for more
// information.
//
// List websites (specified as domain and path) from a project where performance
// data is available.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/page
// insights/pageinsights-list-performance-data-for-project
type ListPerformanceDataForProjectRequest struct {
	ProjectID string
	Domain    *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListPerformanceDataForProjectRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *ListPerformanceDataForProjectRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListPerformanceDataForProjectRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/page-insights", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListPerformanceDataForProjectRequest) query() url.Values {
	q := make(url.Values)
	if r.Domain != nil {
		q.Set("domain", *r.Domain)
	}
	return q
}
