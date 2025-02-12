package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListFeedbackRequest models a request for the 'user-list-feedback' operation. See
// [1] for more information.
//
// Submitted feedback of the given user.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-list-feedback
type ListFeedbackRequest struct {
	UserID  string
	Subject *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListFeedbackRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListFeedbackRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListFeedbackRequest) url() string {
	return fmt.Sprintf("/v2/users/%s/feedback", url.PathEscape(r.UserID))
}

func (r *ListFeedbackRequest) query() url.Values {
	q := make(url.Values)
	if r.Subject != nil {
		q.Set("subject", *r.Subject)
	}
	return q
}
