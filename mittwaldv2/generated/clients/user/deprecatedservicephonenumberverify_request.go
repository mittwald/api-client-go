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

// DeprecatedServicePhoneNumberVerifyRequest models a request for the
// 'deprecated-user-service-phone-number-verify' operation. See [1] for more
// information.
//
// # Verify phone number
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-phone-number-verify
type DeprecatedServicePhoneNumberVerifyRequest struct {
	Body   DeprecatedServicePhoneNumberVerifyRequestBody
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServicePhoneNumberVerifyRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *DeprecatedServicePhoneNumberVerifyRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedServicePhoneNumberVerifyRequest) url() string {
	return fmt.Sprintf("/v2/user/%s/phone/verify", url.PathEscape(r.UserID))
}

func (r *DeprecatedServicePhoneNumberVerifyRequest) query() url.Values {
	return nil
}
