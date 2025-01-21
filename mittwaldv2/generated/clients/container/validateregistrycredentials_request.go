package container

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type ValidateRegistryCredentialsRequest struct {
	RegistryID uuid.UUID
}

func (r *ValidateRegistryCredentialsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ValidateRegistryCredentialsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ValidateRegistryCredentialsRequest) url() string {
	return fmt.Sprintf("/v2/registries/%s/actions/validate-credentials", url.PathEscape(r.RegistryID.String()))
}

func (r *ValidateRegistryCredentialsRequest) query() url.Values {
	return nil
}