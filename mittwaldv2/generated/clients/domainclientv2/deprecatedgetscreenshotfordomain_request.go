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

// DeprecatedGetScreenshotForDomainRequest models a request for the
// 'deprecated-domain-get-screenshot-for-domain' operation. See [1] for more
// information.
//
// Get File Service Reference for a Screenshot of a domain.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/domain/deprecated-domain-get-screenshot-for-domain
type DeprecatedGetScreenshotForDomainRequest struct {
	Body     DeprecatedGetScreenshotForDomainRequestBody
	DomainID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedGetScreenshotForDomainRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *DeprecatedGetScreenshotForDomainRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedGetScreenshotForDomainRequest) url() string {
	return fmt.Sprintf("/v2/domains/%s/screenshots/newest", url.PathEscape(r.DomainID))
}

func (r *DeprecatedGetScreenshotForDomainRequest) query() url.Values {
	return nil
}
