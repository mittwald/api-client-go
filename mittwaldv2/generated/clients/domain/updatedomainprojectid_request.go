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

type UpdateDomainProjectIDRequest struct {
	Body     UpdateDomainProjectIDRequestBody
	DomainID uuid.UUID
}

func (r *UpdateDomainProjectIDRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateDomainProjectIDRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateDomainProjectIDRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/project-id", url.PathEscape(r.DomainID.String()))
}

func (r *UpdateDomainProjectIDRequest) query() url.Values {
	return nil
}
