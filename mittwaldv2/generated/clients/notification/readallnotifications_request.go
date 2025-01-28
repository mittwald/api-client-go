package notification

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

// ReadAllNotificationsRequest models a request for the
// 'notifications-read-all-notifications' operation. See [1] for more information.
//
// Mark all notifications as read.
//
// Mark all notifications for the authenticated user as read.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/notification/notifications-read-all-notifications
type ReadAllNotificationsRequest struct {
	Body               ReadAllNotificationsRequestBody
	Severities         []ReadAllNotificationsRequestQuerySeveritiesItem
	ReferenceID        *string
	ReferenceAggregate *string
	ReferenceDomain    *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ReadAllNotificationsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ReadAllNotificationsRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ReadAllNotificationsRequest) url() string {
	return "/v2/notifications/actions/read-all"
}

func (r *ReadAllNotificationsRequest) query() url.Values {
	q := make(url.Values)
	for _, val := range r.Severities {
		q.Add("severities", string(val))
	}
	if r.ReferenceID != nil {
		q.Set("referenceId", *r.ReferenceID)
	}
	if r.ReferenceAggregate != nil {
		q.Set("referenceAggregate", *r.ReferenceAggregate)
	}
	if r.ReferenceDomain != nil {
		q.Set("referenceDomain", *r.ReferenceDomain)
	}
	return q
}
