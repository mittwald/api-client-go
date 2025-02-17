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

// DeprecatedResendVerificationEmailRequest models a request for the
// 'deprecated-user-resend-verification-email' operation. See [1] for more
// information.
//
// Resend the Email-Address verification email. Replaced by `POST`
// `/v2/users/self/credentials/email/actions/resend-email`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-resend-verification-email
type DeprecatedResendVerificationEmailRequest struct {
	Body DeprecatedResendVerificationEmailRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedResendVerificationEmailRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedResendVerificationEmailRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedResendVerificationEmailRequest) url() string {
	u := url.URL{
		Path: "/v2/signup/email/resend",
	}
	return u.String()
}

func (r *DeprecatedResendVerificationEmailRequest) query() url.Values {
	return nil
}
