package cronjob

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type ListCronjobsRequest struct {
	ProjectID uuid.UUID
	Limit     *int64
	Skip      *int64
	Page      *int64
}

func (r *ListCronjobsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListCronjobsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListCronjobsRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/cronjobs", url.PathEscape(r.ProjectID.String()))
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
