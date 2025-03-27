package contractclientv2

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
func (r *PreviewTariffChangeRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *PreviewTariffChangeRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *PreviewTariffChangeRequest) url() string {
	u := url.URL{
		Path: "/v2/tariff-change-previews",
	}
	return u.String()
}

func (r *PreviewTariffChangeRequest) query() url.Values {
	return nil
}
