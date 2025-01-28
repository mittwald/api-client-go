package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedServicePhoneNumberRemoveRequest models a request for the
// 'deprecated-user-service-phone-number-remove' operation. See [1] for more
// information.
//
// remove your PhoneNumber
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-phone-number-remove
type DeprecatedServicePhoneNumberRemoveRequest struct {
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServicePhoneNumberRemoveRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) url() string {
	return fmt.Sprintf("/v2/user/%s/phone", url.PathEscape(r.UserID))
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) query() url.Values {
	return nil
}
