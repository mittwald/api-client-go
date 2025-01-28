package user

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

// DeprecatedInitPasswordResetRequest models a request for the
// 'deprecated-user-init-password-reset' operation. See [1] for more information.
//
// Initialize password reset process. Replaced by `POST`
// `/v2/users/self/credentials/actions/init-password-reset`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-init-password-reset
type DeprecatedInitPasswordResetRequest struct {
	Body DeprecatedInitPasswordResetRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedInitPasswordResetRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *DeprecatedInitPasswordResetRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedInitPasswordResetRequest) url() string {
	return "/v2/signup/password/reset"
}

func (r *DeprecatedInitPasswordResetRequest) query() url.Values {
	return nil
}
