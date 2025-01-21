package container

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetVolumeRequest struct {
	StackID  uuid.UUID
	VolumeID uuid.UUID
}

func (r *GetVolumeRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetVolumeRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetVolumeRequest) url() string {
	return fmt.Sprintf("/v2/stacks/%s/volumes/%s", url.PathEscape(r.StackID.String()), url.PathEscape(r.VolumeID.String()))
}

func (r *GetVolumeRequest) query() url.Values {
	return nil
}