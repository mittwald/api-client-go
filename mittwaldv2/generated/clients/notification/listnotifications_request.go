package notification

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListNotificationsRequest struct {
	Status *ListNotificationsRequestQueryStatus
	Limit  *int64
	Skip   *int64
	Page   *int64
}

func (r *ListNotificationsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListNotificationsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListNotificationsRequest) url() string {
	return "/v2/notifications"
}

func (r *ListNotificationsRequest) query() url.Values {
	q := make(url.Values)
	if r.Status != nil {
		q.Set("status", string(*r.Status))
	}
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
