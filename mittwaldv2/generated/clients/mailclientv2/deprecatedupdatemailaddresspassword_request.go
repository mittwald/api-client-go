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

// DeprecatedUpdateMailAddressPasswordRequest models a request for the
// 'deprecated-mail-update-mail-address-password' operation. See [1] for more
// information.
//
// Update the password for a MailAddress.
//
// This operation is deprecated. Use the PATCH
// v2/mail-addresses/{mailAddressId}/password endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/deprecated-mail-update-mail-address-password
type DeprecatedUpdateMailAddressPasswordRequest struct {
	Body          DeprecatedUpdateMailAddressPasswordRequestBody
	MailAddressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedUpdateMailAddressPasswordRequest) BuildRequest() (*http.Request, error) {
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

func (r *DeprecatedUpdateMailAddressPasswordRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedUpdateMailAddressPasswordRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/mail-addresses/%s/password", url.PathEscape(r.MailAddressID)),
	}
	return u.String()
}

func (r *DeprecatedUpdateMailAddressPasswordRequest) query() url.Values {
	return nil
}
