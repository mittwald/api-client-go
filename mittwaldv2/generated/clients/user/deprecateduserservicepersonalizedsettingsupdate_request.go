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

type DeprecatedUserServicePersonalizedSettingsUpdateRequest struct {
	Body DeprecatedUserServicePersonalizedSettingsUpdateRequestBody
}

func (r *DeprecatedUserServicePersonalizedSettingsUpdateRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *DeprecatedUserServicePersonalizedSettingsUpdateRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedUserServicePersonalizedSettingsUpdateRequest) url() string {
	return "/v2/user/settings"
}

func (r *DeprecatedUserServicePersonalizedSettingsUpdateRequest) query() url.Values {
	return nil
}
