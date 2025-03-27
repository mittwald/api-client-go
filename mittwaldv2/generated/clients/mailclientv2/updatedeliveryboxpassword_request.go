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

// UpdateDeliveryBoxPasswordRequest models a request for the
// 'mail-update-delivery-box-password' operation. See [1] for more information.
//
// Update the password of a DeliveryBox.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/mail-update-delivery-box-password
type UpdateDeliveryBoxPasswordRequest struct {
	Body          UpdateDeliveryBoxPasswordRequestBody
	DeliveryBoxID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateDeliveryBoxPasswordRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *UpdateDeliveryBoxPasswordRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateDeliveryBoxPasswordRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/delivery-boxes/%s/password", url.PathEscape(r.DeliveryBoxID)),
	}
	return u.String()
}

func (r *UpdateDeliveryBoxPasswordRequest) query() url.Values {
	return nil
}
