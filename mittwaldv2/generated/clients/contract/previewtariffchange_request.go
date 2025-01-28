package contract

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

// PreviewTariffChangeRequest models a request for the
// 'order-preview-tariff-change' operation. See [1] for more information.
//
// Preview TariffChange.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/order-preview-tariff-change
type PreviewTariffChangeRequest struct {
	Body PreviewTariffChangeRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *PreviewTariffChangeRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *PreviewTariffChangeRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *PreviewTariffChangeRequest) url() string {
	return "/v2/tariff-change-previews"
}

func (r *PreviewTariffChangeRequest) query() url.Values {
	return nil
}
