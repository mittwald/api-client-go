package containerclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteVolumeRequest models a request for the 'container-delete-volume'
// operation. See [1] for more information.
//
// Delete a Volume belonging to a Stack.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-delete-volume
type DeleteVolumeRequest struct {
	StackID  string
	VolumeID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteVolumeRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *DeleteVolumeRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteVolumeRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/stacks/%s/volumes/%s", url.PathEscape(r.StackID), url.PathEscape(r.VolumeID)),
	}
	return u.String()
}

func (r *DeleteVolumeRequest) query() url.Values {
	return nil
}
