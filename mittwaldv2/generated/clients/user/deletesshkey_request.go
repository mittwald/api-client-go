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

type DeleteSSHKeyRequest struct {
	SSHKeyID uuid.UUID
}

func (r *DeleteSSHKeyRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteSSHKeyRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteSSHKeyRequest) url() string {
	return fmt.Sprintf("/v2/users/self/ssh-keys/%s", url.PathEscape(r.SSHKeyID.String()))
}

func (r *DeleteSSHKeyRequest) query() url.Values {
	return nil
}