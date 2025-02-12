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

// ConfirmPasswordResetRequest models a request for the
// 'user-confirm-password-reset' operation. See [1] for more information.
//
// Confirm password reset.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-confirm-password-reset
type ConfirmPasswordResetRequest struct {
	Body ConfirmPasswordResetRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ConfirmPasswordResetRequest) BuildRequest() (*http.Request, error) {
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

func (r *ConfirmPasswordResetRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *ConfirmPasswordResetRequest) url() string {
	return "/v2/users/self/credentials/password/confirm-reset"
}

func (r *ConfirmPasswordResetRequest) query() url.Values {
	return nil
}
