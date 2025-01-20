package contract

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

type TerminateContractRequest struct {
	Body       TerminateContractRequestBody
	ContractID uuid.UUID
}

func (r *TerminateContractRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *TerminateContractRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *TerminateContractRequest) url() string {
	return fmt.Sprintf("/v2/contracts/%s/termination", url.PathEscape(r.ContractID.String()))
}

func (r *TerminateContractRequest) query() url.Values {
	return nil
}