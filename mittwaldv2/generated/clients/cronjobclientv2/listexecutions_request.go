package cronjobclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/cronjobv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListExecutionsRequest models a request for the 'cronjob-list-executions'
// operation. See [1] for more information.
//
// List CronjobExecutions belonging to a Cronjob.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/cronjob/cronjob-list-executions
type ListExecutionsRequest struct {
	CronjobID       string
	Limit           *int64
	Skip            *int64
	Page            *int64
	Since           *time.Time
	Until           *time.Time
	Status          *string
	TriggeredByUser *bool
	SortOrder       *cronjobv2.CronjobExecutionSortOrder
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListExecutionsRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListExecutionsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListExecutionsRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/cronjobs/%s/executions", url.PathEscape(r.CronjobID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListExecutionsRequest) query() url.Values {
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
	if r.Since != nil {
		q.Set("since", r.Since.Format(time.RFC3339))
	}
	if r.Until != nil {
		q.Set("until", r.Until.Format(time.RFC3339))
	}
	if r.Status != nil {
		q.Set("status", *r.Status)
	}
	if r.TriggeredByUser != nil {
		q.Set("triggeredByUser", strconv.FormatBool(*r.TriggeredByUser))
	}
	if r.SortOrder != nil {
		q.Set("sortOrder", string(*r.SortOrder))
	}
	return q
}
