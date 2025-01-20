package domain

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

type TLSDeprecatedRequest struct {
	Body      TLSDeprecatedRequestBody
	IngressID uuid.UUID
}

func (r *TLSDeprecatedRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *TLSDeprecatedRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *TLSDeprecatedRequest) url() string {
	return fmt.Sprintf("/v2/ingresses/%s/tls", url.PathEscape(r.IngressID.String()))
}

func (r *TLSDeprecatedRequest) query() url.Values {
	return nil
}
