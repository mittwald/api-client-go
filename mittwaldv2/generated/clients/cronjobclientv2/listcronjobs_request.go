package cronjobclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListCronjobsRequest models a request for the 'cronjob-list-cronjobs' operation.
// See [1] for more information.
//
// List Cronjobs belonging to a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-list-cronjobs
type ListCronjobsRequest struct {
	ProjectID string
	Limit     *int64
	Skip      *int64
	Page      *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListCronjobsRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *ListCronjobsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListCronjobsRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/cronjobs", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListCronjobsRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	return q
}
