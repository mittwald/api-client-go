package domain

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

type ListIngressesCompatibleWithCertificateRequest struct {
	Body ListIngressesCompatibleWithCertificateRequestBody
}

func (r *ListIngressesCompatibleWithCertificateRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *ListIngressesCompatibleWithCertificateRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ListIngressesCompatibleWithCertificateRequest) url() string {
	return "/v2/actions/list-ingresses-compatible-with-certificate"
}

func (r *ListIngressesCompatibleWithCertificateRequest) query() url.Values {
	return nil
}