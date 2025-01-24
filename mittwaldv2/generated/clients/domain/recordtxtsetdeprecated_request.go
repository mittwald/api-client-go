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

type RecordTxtSetDeprecatedRequest struct {
	Body   RecordTxtSetDeprecatedRequestBody
	ZoneID uuid.UUID
}

func (r *RecordTxtSetDeprecatedRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *RecordTxtSetDeprecatedRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *RecordTxtSetDeprecatedRequest) url() string {
	return fmt.Sprintf("/v2/dns/zones/%s/recordset/txt", url.PathEscape(r.ZoneID.String()))
}

func (r *RecordTxtSetDeprecatedRequest) query() url.Values {
	return nil
}
