package userclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedServicePersonalizedSettingsGetRequest models a request for the
// 'deprecated-user-service-personalized-settings-get' operation. See [1] for more
// information.
//
// # Get personalized settings for the user executing the request
//
// Use /v2/users/{userId}/settings instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-personalized-settings-get
type DeprecatedServicePersonalizedSettingsGetRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServicePersonalizedSettingsGetRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *DeprecatedServicePersonalizedSettingsGetRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedServicePersonalizedSettingsGetRequest) url() string {
	u := url.URL{
		Path: "/v2/user/settings",
	}
	return u.String()
}

func (r *DeprecatedServicePersonalizedSettingsGetRequest) query() url.Values {
	return nil
}
