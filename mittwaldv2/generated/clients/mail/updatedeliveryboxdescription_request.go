package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type UpdateDeliveryBoxDescriptionRequest struct {
	Body          UpdateDeliveryBoxDescriptionRequestBody
	DeliveryBoxID uuid.UUID
}

func (r *UpdateDeliveryBoxDescriptionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateDeliveryBoxDescriptionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateDeliveryBoxDescriptionRequest) url() string {
	return fmt.Sprintf("/v2/delivery-boxes/%s/description", url.PathEscape(r.DeliveryBoxID.String()))
}

func (r *UpdateDeliveryBoxDescriptionRequest) query() url.Values {
	return nil
}
