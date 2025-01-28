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

// AuthenticateWithAccessTokenRetrievalKeyRequest models a request for the
// 'user-authenticate-with-access-token-retrieval-key' operation. See [1] for more
// information.
//
// Authenticate an user with an access token retrieval key.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-authenticate-with-access-token-retrieval-key
type AuthenticateWithAccessTokenRetrievalKeyRequest struct {
	Body AuthenticateWithAccessTokenRetrievalKeyRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *AuthenticateWithAccessTokenRetrievalKeyRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *AuthenticateWithAccessTokenRetrievalKeyRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *AuthenticateWithAccessTokenRetrievalKeyRequest) url() string {
	return "/v2/authenticate-token-retrieval-key"
}

func (r *AuthenticateWithAccessTokenRetrievalKeyRequest) query() url.Values {
	return nil
}
