package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetPollStatusRequest models a request for the 'user-get-poll-status' operation.
// See [1] for more information.
//
// Get poll settings for the specified user.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-get-poll-status
type GetPollStatusRequest struct {
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetPollStatusRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetPollStatusRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetPollStatusRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/poll-settings/%s", url.PathEscape(r.UserID)),
	}
	return u.String()
}

func (r *GetPollStatusRequest) query() url.Values {
	return nil
}
