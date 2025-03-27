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
func (r *DeprecatedInitPasswordResetRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *DeprecatedInitPasswordResetRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedInitPasswordResetRequest) url() string {
	u := url.URL{
		Path: "/v2/signup/password/reset",
	}
	return u.String()
}

func (r *DeprecatedInitPasswordResetRequest) query() url.Values {
	return nil
}
