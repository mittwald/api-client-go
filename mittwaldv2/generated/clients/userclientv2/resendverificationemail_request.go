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
func (r *ResendVerificationEmailRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *ResendVerificationEmailRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *ResendVerificationEmailRequest) url() string {
	u := url.URL{
		Path: "/v2/users/self/credentials/email/actions/resend-email",
	}
	return u.String()
}

func (r *ResendVerificationEmailRequest) query() url.Values {
	return nil
}
