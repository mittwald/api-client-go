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
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *AuthenticateRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *AuthenticateRequest) url() string {
	return "/v2/authenticate"
}

func (r *AuthenticateRequest) query() url.Values {
	return nil
}
