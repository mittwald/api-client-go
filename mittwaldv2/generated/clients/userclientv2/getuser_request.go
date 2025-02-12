package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetUserRequest models a request for the 'user-get-user' operation. See [1] for
// more information.
//
// Get profile information for a user.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-get-user
type GetUserRequest struct {
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetUserRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetUserRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetUserRequest) url() string {
	return fmt.Sprintf("/v2/users/%s", url.PathEscape(r.UserID))
}

func (r *GetUserRequest) query() url.Values {
	return nil
}
