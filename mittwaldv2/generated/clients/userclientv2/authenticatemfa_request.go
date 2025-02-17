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

// AuthenticateMFARequest models a request for the 'user-authenticate-mfa'
// operation. See [1] for more information.
//
// Validate your second factor.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-authenticate-mfa
type AuthenticateMFARequest struct {
	Body AuthenticateMFARequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AuthenticateMFARequest) BuildRequest() (*http.Request, error) {
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

func (r *AuthenticateMFARequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *AuthenticateMFARequest) url() string {
	u := url.URL{
		Path: "/v2/authenticate-mfa",
	}
	return u.String()
}

func (r *AuthenticateMFARequest) query() url.Values {
	return nil
}
