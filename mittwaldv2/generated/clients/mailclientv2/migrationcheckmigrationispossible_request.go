package mailclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// MigrationCheckMigrationIsPossibleRequest models a request for the
// 'mail-migration-check-migration-is-possible' operation. See [1] for more
// information.
//
// Check if a Migration between two projects is possible.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/mail-migration-check-migration-is-possible
type MigrationCheckMigrationIsPossibleRequest struct {
	Body MigrationCheckMigrationIsPossibleRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *MigrationCheckMigrationIsPossibleRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *MigrationCheckMigrationIsPossibleRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *MigrationCheckMigrationIsPossibleRequest) url() string {
	u := url.URL{
		Path: "/v2/mail-migrations/actions/possibility-check",
	}
	return u.String()
}

func (r *MigrationCheckMigrationIsPossibleRequest) query() url.Values {
	return nil
}
