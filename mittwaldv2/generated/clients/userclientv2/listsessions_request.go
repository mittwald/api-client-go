package userclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListSessionsRequest models a request for the 'user-list-sessions' operation. See
// [1] for more information.
//
// List all sessions.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-list-sessions
type ListSessionsRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListSessionsRequest) BuildRequest() (*http.Request, error) {
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

func (r *ListSessionsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListSessionsRequest) url() string {
	u := url.URL{
		Path: "/v2/users/self/sessions",
	}
	return u.String()
}

func (r *ListSessionsRequest) query() url.Values {
	return nil
}
