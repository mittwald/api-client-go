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

// AuthenticateRequest models a request for the 'user-authenticate' operation. See
// [1] for more information.
//
// Authenticate yourself to get an access token.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-authenticate
type AuthenticateRequest struct {
	Body AuthenticateRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AuthenticateRequest) BuildRequest() (*http.Request, error) {
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

func (r *AuthenticateRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *AuthenticateRequest) url() string {
	u := url.URL{
		Path: "/v2/authenticate",
	}
	return u.String()
}

func (r *AuthenticateRequest) query() url.Values {
	return nil
}
