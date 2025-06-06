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

// CreateDNSZoneRequest models a request for the 'dns-create-dns-zone' operation.
// See [1] for more information.
//
// Create a DNSZone.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/domain/dns-create-dns-zone
type CreateDNSZoneRequest struct {
	Body CreateDNSZoneRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateDNSZoneRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *CreateDNSZoneRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateDNSZoneRequest) url() string {
	u := url.URL{
		Path: "/v2/dns-zones",
	}
	return u.String()
}

func (r *CreateDNSZoneRequest) query() url.Values {
	return nil
}
