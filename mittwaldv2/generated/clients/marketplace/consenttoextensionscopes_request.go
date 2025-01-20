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

type ConsentToExtensionScopesRequest struct {
	Body                ConsentToExtensionScopesRequestBody
	ExtensionInstanceID uuid.UUID
}

func (r *ConsentToExtensionScopesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *ConsentToExtensionScopesRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ConsentToExtensionScopesRequest) url() string {
	return fmt.Sprintf("/v2/extension-instances/%s/scopes", url.PathEscape(r.ExtensionInstanceID.String()))
}

func (r *ConsentToExtensionScopesRequest) query() url.Values {
	return nil
}