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

// DeprecatedUpdateMailAddressCatchallRequest models a request for the
// 'deprecated-mail-update-mail-address-catchall' operation. See [1] for more
// information.
//
// Update the catchall of a MailAddress.
//
// This operation is deprecated. Use the PATCH
// v2/mail-addresses/{mailAddressId}/catch-all endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/deprecated-mail-update-mail-address-catchall
type DeprecatedUpdateMailAddressCatchallRequest struct {
	Body          DeprecatedUpdateMailAddressCatchallRequestBody
	MailAddressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedUpdateMailAddressCatchallRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *DeprecatedUpdateMailAddressCatchallRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedUpdateMailAddressCatchallRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/mail-addresses/%s/catchall", url.PathEscape(r.MailAddressID)),
	}
	return u.String()
}

func (r *DeprecatedUpdateMailAddressCatchallRequest) query() url.Values {
	return nil
}
