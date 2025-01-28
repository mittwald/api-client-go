package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// PasswordValidationGetPasswordPolicyRequest models a request for the
// 'password-validation-get-password-policy' operation. See [1] for more
// information.
//
// Get a PasswordPolicy.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/password-validation-get-password-policy
type PasswordValidationGetPasswordPolicyRequest struct {
	PasswordPolicy string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *PasswordValidationGetPasswordPolicyRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *PasswordValidationGetPasswordPolicyRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *PasswordValidationGetPasswordPolicyRequest) url() string {
	return fmt.Sprintf("/v2/password-policies/%s", url.PathEscape(r.PasswordPolicy))
}

func (r *PasswordValidationGetPasswordPolicyRequest) query() url.Values {
	return nil
}
