package marketplace

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

type RotateSecretForExtensionInstanceRequest struct {
	Body                RotateSecretForExtensionInstanceRequestBody
	ContributorID       string
	ExtensionID         uuid.UUID
	ExtensionInstanceID uuid.UUID
}

func (r *RotateSecretForExtensionInstanceRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *RotateSecretForExtensionInstanceRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *RotateSecretForExtensionInstanceRequest) url() string {
	return fmt.Sprintf("/v2/contributors/%s/extensions/%s/extension-instances/%s/secret", url.PathEscape(r.ContributorID), url.PathEscape(r.ExtensionID.String()), url.PathEscape(r.ExtensionInstanceID.String()))
}

func (r *RotateSecretForExtensionInstanceRequest) query() url.Values {
	return nil
}
