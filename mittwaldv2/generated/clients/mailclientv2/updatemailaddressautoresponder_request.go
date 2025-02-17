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

// UpdateMailAddressAutoresponderRequest models a request for the
// 'mail-update-mail-address-autoresponder' operation. See [1] for more
// information.
//
// Update the autoresponder of a MailAddress.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/mail-update-mail-address-autoresponder
type UpdateMailAddressAutoresponderRequest struct {
	Body          UpdateMailAddressAutoresponderRequestBody
	MailAddressID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateMailAddressAutoresponderRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *UpdateMailAddressAutoresponderRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateMailAddressAutoresponderRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/mail-addresses/%s/autoresponder", url.PathEscape(r.MailAddressID)),
	}
	return u.String()
}

func (r *UpdateMailAddressAutoresponderRequest) query() url.Values {
	return nil
}
