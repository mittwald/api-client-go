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

// CreateAPITokenRequest models a request for the 'user-create-api-token'
// operation. See [1] for more information.
//
// Store a new ApiToken.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-create-api-token
type CreateAPITokenRequest struct {
	Body CreateAPITokenRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateAPITokenRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *CreateAPITokenRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateAPITokenRequest) url() string {
	u := url.URL{
		Path: "/v2/users/self/api-tokens",
	}
	return u.String()
}

func (r *CreateAPITokenRequest) query() url.Values {
	return nil
}
