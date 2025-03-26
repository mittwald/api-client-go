package marketplaceclientv2

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

// AuthenticateWithSessionTokenRequest models a request for the
// 'extension-authenticate-with-session-token' operation. See [1] for more
// information.
//
// Authenticate your external application using a session token and an extension
// secret
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-authenticate-with-session-token
type AuthenticateWithSessionTokenRequest struct {
	Body AuthenticateWithSessionTokenRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AuthenticateWithSessionTokenRequest) BuildRequest() (*http.Request, error) {
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

func (r *AuthenticateWithSessionTokenRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *AuthenticateWithSessionTokenRequest) url() string {
	u := url.URL{
		Path: "/v2/authenticate-session-token",
	}
	return u.String()
}

func (r *AuthenticateWithSessionTokenRequest) query() url.Values {
	return nil
}
