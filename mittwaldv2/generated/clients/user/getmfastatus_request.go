package user

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetMFAStatusRequest models a request for the 'user-get-mfa-status' operation.
// See [1] for more information.
//
// Get your current multi factor auth status.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-get-mfa-status
type GetMFAStatusRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetMFAStatusRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetMFAStatusRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetMFAStatusRequest) url() string {
	return "/v2/users/self/credentials/mfa"
}

func (r *GetMFAStatusRequest) query() url.Values {
	return nil
}
