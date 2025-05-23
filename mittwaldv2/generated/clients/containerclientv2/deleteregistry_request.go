package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteRegistryRequest models a request for the 'container-delete-registry'
// operation. See [1] for more information.
//
// Delete a Registry.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-delete-registry
type DeleteRegistryRequest struct {
	RegistryID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteRegistryRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
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

func (r *DeleteRegistryRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteRegistryRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/registries/%s", url.PathEscape(r.RegistryID)),
	}
	return u.String()
}

func (r *DeleteRegistryRequest) query() url.Values {
	return nil
}
