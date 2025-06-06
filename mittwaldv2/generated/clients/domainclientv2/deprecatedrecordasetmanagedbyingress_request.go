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

// DeprecatedRecordASetManagedByIngressRequest models a request for the
// 'deprecated-dns-record-a-set-managed-by-ingress' operation. See [1] for more
// information.
//
// set a-records managed by ingress for a specific zone
//
// This operation is deprecated. Use the POST
// v2/dns-zones/{dnsZoneId}/record-sets/{recordSet}/actions/set-managed endpoint
// instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/deprecated-dns-record-a-set-managed-by-ingress
type DeprecatedRecordASetManagedByIngressRequest struct {
	Body   DeprecatedRecordASetManagedByIngressRequestBody
	ZoneID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedRecordASetManagedByIngressRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *DeprecatedRecordASetManagedByIngressRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedRecordASetManagedByIngressRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/dns/zones/%s/recordset/acombined/managed/ingress", url.PathEscape(r.ZoneID)),
	}
	return u.String()
}

func (r *DeprecatedRecordASetManagedByIngressRequest) query() url.Values {
	return nil
}
