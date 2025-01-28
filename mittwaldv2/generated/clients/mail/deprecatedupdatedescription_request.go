package mail

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

// DeprecatedUpdateDescriptionRequest models a request for the
// 'deprecated-mail-deliverybox-update-description' operation. See [1] for more
// information.
//
// # Update the description of an deliverybox
//
// This operation is deprecated. Use the PATCH
// v2/delivery-boxes/{deliveryBoxId}/description endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/deprecated-mail-deliverybox-update-description
type DeprecatedUpdateDescriptionRequest struct {
	Body DeprecatedUpdateDescriptionRequestBody
	ID   string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedUpdateDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *DeprecatedUpdateDescriptionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedUpdateDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/deliveryboxes/%s/description", url.PathEscape(r.ID))
}

func (r *DeprecatedUpdateDescriptionRequest) query() url.Values {
	return nil
}
