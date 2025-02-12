package userclientv2

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

// DeprecatedConfirmPasswordResetRequest models a request for the
// 'deprecated-user-confirm-password-reset' operation. See [1] for more
// information.
//
// Confirm password reset. Replaced by `POST`
// `/v2/users/self/credentials/password/confirm-reset`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-confirm-password-reset
type DeprecatedConfirmPasswordResetRequest struct {
	Body DeprecatedConfirmPasswordResetRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedConfirmPasswordResetRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedConfirmPasswordResetRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedConfirmPasswordResetRequest) url() string {
	return "/v2/signup/password/reset/confirm"
}

func (r *DeprecatedConfirmPasswordResetRequest) query() url.Values {
	return nil
}
