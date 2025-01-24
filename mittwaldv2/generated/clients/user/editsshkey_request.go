package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type EditSSHKeyRequest struct {
	Body     EditSSHKeyRequestBody
	SSHKeyID uuid.UUID
}

func (r *EditSSHKeyRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *EditSSHKeyRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *EditSSHKeyRequest) url() string {
	return fmt.Sprintf("/v2/users/self/ssh-keys/%s", url.PathEscape(r.SSHKeyID.String()))
}

func (r *EditSSHKeyRequest) query() url.Values {
	return nil
}
