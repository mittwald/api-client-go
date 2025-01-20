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

type UpdateMysqlDatabaseDescriptionRequest struct {
	Body            UpdateMysqlDatabaseDescriptionRequestBody
	MysqlDatabaseID uuid.UUID
}

func (r *UpdateMysqlDatabaseDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateMysqlDatabaseDescriptionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateMysqlDatabaseDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/mysql-databases/%s/description", url.PathEscape(r.MysqlDatabaseID.String()))
}

func (r *UpdateMysqlDatabaseDescriptionRequest) query() url.Values {
	return nil
}