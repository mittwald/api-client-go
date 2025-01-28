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

// VerifyPhoneNumberRequest models a request for the 'user-verify-phone-number'
// operation. See [1] for more information.
//
// Verify phone number.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-verify-phone-number
type VerifyPhoneNumberRequest struct {
	Body   VerifyPhoneNumberRequestBody
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *VerifyPhoneNumberRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *VerifyPhoneNumberRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *VerifyPhoneNumberRequest) url() string {
	return fmt.Sprintf("/v2/users/%s/actions/verify-phone", url.PathEscape(r.UserID))
}

func (r *VerifyPhoneNumberRequest) query() url.Values {
	return nil
}
