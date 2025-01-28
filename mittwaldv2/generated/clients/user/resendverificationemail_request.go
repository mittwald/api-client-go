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

// ResendVerificationEmailRequest models a request for the
// 'user-resend-verification-email' operation. See [1] for more information.
//
// Resend the Email-Address verification email.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-resend-verification-email
type ResendVerificationEmailRequest struct {
	Body ResendVerificationEmailRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ResendVerificationEmailRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ResendVerificationEmailRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ResendVerificationEmailRequest) url() string {
	return "/v2/users/self/credentials/email/actions/resend-email"
}

func (r *ResendVerificationEmailRequest) query() url.Values {
	return nil
}
