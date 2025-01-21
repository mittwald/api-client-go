package container

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

type GetContainerImageConfigRequest struct {
	ImageReference              string
	UseCredentialsForProjectID  *uuid.UUID
	UseCredentialsForRegistryID *uuid.UUID
	GenerateAiData              *bool
}

func (r *GetContainerImageConfigRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetContainerImageConfigRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetContainerImageConfigRequest) url() string {
	return "/v2/container-image-config"
}

func (r *GetContainerImageConfigRequest) query() url.Values {
	q := make(url.Values)
	q.Set("imageReference", r.ImageReference)
	if r.UseCredentialsForProjectID != nil {
		q.Set("useCredentialsForProjectId", r.UseCredentialsForProjectID.String())
	}
	if r.UseCredentialsForRegistryID != nil {
		q.Set("useCredentialsForRegistryId", r.UseCredentialsForRegistryID.String())
	}
	if r.GenerateAiData != nil {
		q.Set("generateAiData", strconv.FormatBool(*r.GenerateAiData))
	}
	return q
}
