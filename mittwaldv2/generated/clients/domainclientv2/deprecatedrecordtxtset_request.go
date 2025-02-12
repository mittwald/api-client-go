package domainclientv2

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

// DeprecatedRecordTxtSetRequest models a request for the
// 'deprecated-dns-record-txt-set' operation. See [1] for more information.
//
// updates txt-records for a specific zone
//
// This operation is deprecated. Use the PUT
// v2/dns-zones/{dnsZoneId}/record-sets/{recordSet} endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/deprecated-dns-record-txt-set
type DeprecatedRecordTxtSetRequest struct {
	Body   DeprecatedRecordTxtSetRequestBody
	ZoneID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedRecordTxtSetRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedRecordTxtSetRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedRecordTxtSetRequest) url() string {
	return fmt.Sprintf("/v2/dns/zones/%s/recordset/txt", url.PathEscape(r.ZoneID))
}

func (r *DeprecatedRecordTxtSetRequest) query() url.Values {
	return nil
}
