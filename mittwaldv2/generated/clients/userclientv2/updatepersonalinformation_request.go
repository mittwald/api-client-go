package userclientv2

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

// UpdatePersonalInformationRequest models a request for the
// 'user-update-personal-information' operation. See [1] for more information.
//
// Change personal information.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-update-personal-information
type UpdatePersonalInformationRequest struct {
	Body   UpdatePersonalInformationRequestBody
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdatePersonalInformationRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *UpdatePersonalInformationRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdatePersonalInformationRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/users/%s", url.PathEscape(r.UserID)),
	}
	return u.String()
}

func (r *UpdatePersonalInformationRequest) query() url.Values {
	return nil
}
