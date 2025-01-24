package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type MigrationGetMigrationRequest struct {
	MigrationID uuid.UUID
}

func (r *MigrationGetMigrationRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *MigrationGetMigrationRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *MigrationGetMigrationRequest) url() string {
	return fmt.Sprintf("/v2/mail-migrations/%s", url.PathEscape(r.MigrationID.String()))
}

func (r *MigrationGetMigrationRequest) query() url.Values {
	return nil
}