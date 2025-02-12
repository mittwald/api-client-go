package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedServiceUserGetRequest models a request for the
// 'deprecated-user-service-user-get' operation. See [1] for more information.
//
// Get profile information for the specified user if the user is related to the
// executing user
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-user-get
type DeprecatedServiceUserGetRequest struct {
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServiceUserGetRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedServiceUserGetRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedServiceUserGetRequest) url() string {
	return fmt.Sprintf("/v2/user/%s", url.PathEscape(r.UserID))
}

func (r *DeprecatedServiceUserGetRequest) query() url.Values {
	return nil
}
