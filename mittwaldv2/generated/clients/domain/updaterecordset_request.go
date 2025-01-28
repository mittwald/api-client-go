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

// UpdateRecordSetRequest models a request for the 'dns-update-record-set'
// operation. See [1] for more information.
//
// Update a record set on a DNSZone.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/dns-update-record-set
type UpdateRecordSetRequest struct {
	Body      UpdateRecordSetRequestBody
	DNSZoneID string
	RecordSet UpdateRecordSetRequestPathRecordSet
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateRecordSetRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *UpdateRecordSetRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateRecordSetRequest) url() string {
	return fmt.Sprintf("/v2/dns-zones/%s/record-sets/%s", url.PathEscape(r.DNSZoneID), url.PathEscape(string(r.RecordSet)))
}

func (r *UpdateRecordSetRequest) query() url.Values {
	return nil
}
