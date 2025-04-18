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

// DeprecatedEditAPITokenRequest models a request for the
// 'deprecated-user-edit-api-token' operation. See [1] for more information.
//
// Update an existing `ApiToken`. Replaced by `PUT`
// `/v2/users/{userId}/api-tokens/{apiTokenId}`.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-edit-api-token
type DeprecatedEditAPITokenRequest struct {
	Body       DeprecatedEditAPITokenRequestBody
	APITokenID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedEditAPITokenRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *DeprecatedEditAPITokenRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedEditAPITokenRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/signup/token/api/%s", url.PathEscape(r.APITokenID)),
	}
	return u.String()
}

func (r *DeprecatedEditAPITokenRequest) query() url.Values {
	return nil
}
