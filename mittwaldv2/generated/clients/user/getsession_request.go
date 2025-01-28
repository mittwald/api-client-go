package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetSessionRequest models a request for the 'user-get-session' operation. See [1]
// for more information.
//
// Get a specific session.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-get-session
type GetSessionRequest struct {
	TokenID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetSessionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetSessionRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetSessionRequest) url() string {
	return fmt.Sprintf("/v2/users/self/sessions/%s", url.PathEscape(r.TokenID))
}

func (r *GetSessionRequest) query() url.Values {
	return nil
}
