package userclientv2

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
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/user/%s/phone", url.PathEscape(r.UserID)),
	}
	return u.String()
}

func (r *DeprecatedServicePhoneNumberRemoveRequest) query() url.Values {
	return nil
}
