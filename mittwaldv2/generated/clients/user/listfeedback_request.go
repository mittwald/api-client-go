package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ListFeedbackRequest struct {
	UserID  string
	Subject *string
}

func (r *ListFeedbackRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListFeedbackRequest) body() (io.Reader, error) {
	return nil, nil
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