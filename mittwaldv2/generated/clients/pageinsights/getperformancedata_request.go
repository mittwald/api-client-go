package pageinsights

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetPerformanceDataRequest models a request for the
// 'pageinsights-get-performance-data' operation. See [1] for more information.
//
// Get detailed performance data for a given domain and path.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/page
// insights/pageinsights-get-performance-data
type GetPerformanceDataRequest struct {
	Domain string
	Path   string
	Date   *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetPerformanceDataRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetPerformanceDataRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetPerformanceDataRequest) url() string {
	return "/v2/page-insights"
}

func (r *GetPerformanceDataRequest) query() url.Values {
	q := make(url.Values)
	q.Set("domain", r.Domain)
	q.Set("path", r.Path)
	if r.Date != nil {
		q.Set("date", *r.Date)
	}
	return q
}
