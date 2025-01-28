package domain

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

// SetRecordSetManagedRequest models a request for the 'dns-set-record-set-managed'
// operation. See [1] for more information.
//
// Set a record set on a DNSZone to managed.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/dns-set-record-set-managed
type SetRecordSetManagedRequest struct {
	Body      SetRecordSetManagedRequestBody
	DNSZoneID string
	RecordSet SetRecordSetManagedRequestPathRecordSet
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *SetRecordSetManagedRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *SetRecordSetManagedRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *SetRecordSetManagedRequest) url() string {
	return fmt.Sprintf("/v2/dns-zones/%s/record-sets/%s/actions/set-managed", url.PathEscape(r.DNSZoneID), url.PathEscape(string(r.RecordSet)))
}

func (r *SetRecordSetManagedRequest) query() url.Values {
	return nil
}
