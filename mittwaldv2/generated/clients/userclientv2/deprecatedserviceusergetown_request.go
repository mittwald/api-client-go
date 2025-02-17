package userclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedServiceUserGetOwnRequest models a request for the
// 'deprecated-user-service-user-get-own' operation. See [1] for more information.
//
// # Get profile information for the executing user
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-user-get-own
type DeprecatedServiceUserGetOwnRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServiceUserGetOwnRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedServiceUserGetOwnRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedServiceUserGetOwnRequest) url() string {
	u := url.URL{
		Path: "/v2/user",
	}
	return u.String()
}

func (r *DeprecatedServiceUserGetOwnRequest) query() url.Values {
	return nil
}
