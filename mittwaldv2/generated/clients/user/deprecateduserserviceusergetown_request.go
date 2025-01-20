package user

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeprecatedUserServiceUserGetOwnRequest struct {
}

func (r *DeprecatedUserServiceUserGetOwnRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *DeprecatedUserServiceUserGetOwnRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedUserServiceUserGetOwnRequest) url() string {
	return "/v2/user"
}

func (r *DeprecatedUserServiceUserGetOwnRequest) query() url.Values {
	return nil
}
