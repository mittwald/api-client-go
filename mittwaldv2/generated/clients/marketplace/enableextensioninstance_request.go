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

type EnableExtensionInstanceRequest struct {
	ExtensionInstanceID uuid.UUID
}

func (r *EnableExtensionInstanceRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *EnableExtensionInstanceRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *EnableExtensionInstanceRequest) url() string {
	return fmt.Sprintf("/v2/extension-instances/%s/actions/enable", url.PathEscape(r.ExtensionInstanceID.String()))
}

func (r *EnableExtensionInstanceRequest) query() url.Values {
	return nil
}
