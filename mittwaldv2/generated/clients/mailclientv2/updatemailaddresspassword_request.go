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

// UpdateMailAddressPasswordRequest models a request for the
// 'mail-update-mail-address-password' operation. See [1] for more information.
//
// Update the password for a MailAddress.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/mail-update-mail-address-password
type UpdateMailAddressPasswordRequest struct {
	Body          UpdateMailAddressPasswordRequestBody
	MailAddressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateMailAddressPasswordRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
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

func (r *UpdateMailAddressPasswordRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateMailAddressPasswordRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/mail-addresses/%s/password", url.PathEscape(r.MailAddressID)),
	}
	return u.String()
}

func (r *UpdateMailAddressPasswordRequest) query() url.Values {
	return nil
}
