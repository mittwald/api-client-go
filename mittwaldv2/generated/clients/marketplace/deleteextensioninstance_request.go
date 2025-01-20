package marketplace

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type DeleteExtensionInstanceRequest struct {
	ExtensionInstanceID uuid.UUID
}

func (r *DeleteExtensionInstanceRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, r.url(), body)
}

func (r *DeleteExtensionInstanceRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeleteExtensionInstanceRequest) url() string {
	return fmt.Sprintf("/v2/extension-instances/%s", url.PathEscape(r.ExtensionInstanceID.String()))
}

func (r *DeleteExtensionInstanceRequest) query() url.Values {
	return nil
}