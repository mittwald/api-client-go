package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeprecatedServiceUserGetRequest struct {
	UserID string
}

func (r *DeprecatedServiceUserGetRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *DeprecatedServiceUserGetRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedServiceUserGetRequest) url() string {
	return fmt.Sprintf("/v2/user/%s", url.PathEscape(r.UserID))
}

func (r *DeprecatedServiceUserGetRequest) query() url.Values {
	return nil
}
