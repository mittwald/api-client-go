package database

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

type UpdateMysqlDatabaseDefaultCharsetRequest struct {
	Body            UpdateMysqlDatabaseDefaultCharsetRequestBody
	MysqlDatabaseID uuid.UUID
}

func (r *UpdateMysqlDatabaseDefaultCharsetRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateMysqlDatabaseDefaultCharsetRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateMysqlDatabaseDefaultCharsetRequest) url() string {
	return fmt.Sprintf("/v2/mysql-databases/%s/default-charset", url.PathEscape(r.MysqlDatabaseID.String()))
}

func (r *UpdateMysqlDatabaseDefaultCharsetRequest) query() url.Values {
	return nil
}