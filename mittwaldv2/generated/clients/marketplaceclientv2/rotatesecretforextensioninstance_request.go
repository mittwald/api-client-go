package marketplaceclientv2

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

// RotateSecretForExtensionInstanceRequest models a request for the
// 'contributor-rotate-secret-for-extension-instance' operation. See [1] for more
// information.
//
// Rotate the secret for an extension instance.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/contributor-rotate-secret-for-extension-instance
type RotateSecretForExtensionInstanceRequest struct {
	Body                RotateSecretForExtensionInstanceRequestBody
	ContributorID       string
	ExtensionID         string
	ExtensionInstanceID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *RotateSecretForExtensionInstanceRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *RotateSecretForExtensionInstanceRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *RotateSecretForExtensionInstanceRequest) url() string {
	return fmt.Sprintf("/v2/contributors/%s/extensions/%s/extension-instances/%s/secret", url.PathEscape(r.ContributorID), url.PathEscape(r.ExtensionID), url.PathEscape(r.ExtensionInstanceID))
}

func (r *RotateSecretForExtensionInstanceRequest) query() url.Values {
	return nil
}
