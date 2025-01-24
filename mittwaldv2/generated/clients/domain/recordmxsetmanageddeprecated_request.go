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

type RecordMxSetManagedDeprecatedRequest struct {
	Body   RecordMxSetManagedDeprecatedRequestBody
	ZoneID uuid.UUID
}

func (r *RecordMxSetManagedDeprecatedRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *RecordMxSetManagedDeprecatedRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *RecordMxSetManagedDeprecatedRequest) url() string {
	return fmt.Sprintf("/v2/dns/zones/%s/recordset/mx/managed", url.PathEscape(r.ZoneID.String()))
}

func (r *RecordMxSetManagedDeprecatedRequest) query() url.Values {
	return nil
}
