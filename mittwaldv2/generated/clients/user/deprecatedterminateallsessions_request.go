package user

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedTerminateAllSessionsRequest models a request for the
// 'deprecated-user-terminate-all-sessions' operation. See [1] for more
// information.
//
// Terminate all sessions, except the current session. Replaced by `DELETE`
// `/v2/users/self/sessions`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-terminate-all-sessions
type DeprecatedTerminateAllSessionsRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedTerminateAllSessionsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeprecatedTerminateAllSessionsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedTerminateAllSessionsRequest) url() string {
	return "/v2/signup/sessions"
}

func (r *DeprecatedTerminateAllSessionsRequest) query() url.Values {
	return nil
}
