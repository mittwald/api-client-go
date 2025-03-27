package notificationclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// CountUnreadNotificationsRequest models a request for the
// 'notifications-count-unread-notifications' operation. See [1] for more
// information.
//
// Get the counts for unread notifications of the user.
//
// The user is determined by the access token used.
// Lighter weight route instead of getting all unread notifications and count
// yourself.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/notification/notifications-count-unread-notifications
type CountUnreadNotificationsRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CountUnreadNotificationsRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *CountUnreadNotificationsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *CountUnreadNotificationsRequest) url() string {
	u := url.URL{
		Path: "/v2/notification-unread-counts",
	}
	return u.String()
}

func (r *CountUnreadNotificationsRequest) query() url.Values {
	return nil
}
