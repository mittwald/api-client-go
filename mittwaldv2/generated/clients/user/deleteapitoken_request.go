package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type DeleteAPITokenRequest struct {
	APITokenID uuid.UUID
}

func (r *DeleteAPITokenRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteAPITokenRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteAPITokenRequest) url() string {
	return fmt.Sprintf("/v2/users/self/api-tokens/%s", url.PathEscape(r.APITokenID.String()))
}

func (r *DeleteAPITokenRequest) query() url.Values {
	return nil
}
