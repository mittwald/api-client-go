package user

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

type VerifyEmailRequest struct {
	Body VerifyEmailRequestBody
}

func (r *VerifyEmailRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *VerifyEmailRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *VerifyEmailRequest) url() string {
	return "/v2/users/self/credentials/email/actions/verify-email"
}

func (r *VerifyEmailRequest) query() url.Values {
	return nil
}