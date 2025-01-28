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

// EditAPITokenRequest models a request for the 'user-edit-api-token' operation.
// See [1] for more information.
//
// Update an existing `ApiToken`.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-edit-api-token
type EditAPITokenRequest struct {
	Body       EditAPITokenRequestBody
	APITokenID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *EditAPITokenRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *EditAPITokenRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *EditAPITokenRequest) url() string {
	return fmt.Sprintf("/v2/users/self/api-tokens/%s", url.PathEscape(r.APITokenID))
}

func (r *EditAPITokenRequest) query() url.Values {
	return nil
}
