package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeprecatedUserTerminateSessionRequest struct {
	TokenID string
}

func (r *DeprecatedUserTerminateSessionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeprecatedUserTerminateSessionRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedUserTerminateSessionRequest) url() string {
	return fmt.Sprintf("/v2/signup/sessions/%s", url.PathEscape(r.TokenID))
}

func (r *DeprecatedUserTerminateSessionRequest) query() url.Values {
	return nil
}