package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeprecatedServiceAvatarRemoveRequest models a request for the
// 'deprecated-user-service-avatar-remove' operation. See [1] for more information.
//
// # Remove Avatar
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-avatar-remove
type DeprecatedServiceAvatarRemoveRequest struct {
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServiceAvatarRemoveRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedServiceAvatarRemoveRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeprecatedServiceAvatarRemoveRequest) url() string {
	return fmt.Sprintf("/v2/user/%s/avatar", url.PathEscape(r.UserID))
}

func (r *DeprecatedServiceAvatarRemoveRequest) query() url.Values {
	return nil
}
