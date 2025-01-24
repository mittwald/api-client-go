package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type RemoveAvatarRequest struct {
	UserID string
}

func (r *RemoveAvatarRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *RemoveAvatarRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *RemoveAvatarRequest) url() string {
	return fmt.Sprintf("/v2/users/%s/avatar", url.PathEscape(r.UserID))
}

func (r *RemoveAvatarRequest) query() url.Values {
	return nil
}
