package container

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

// UpdateStackRequest models a request for the 'container-update-stack' operation.
// See [1] for more information.
//
// Create, update or delete Services or Volumes belonging to a Stack.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/container/container-update-stack
type UpdateStackRequest struct {
	Body    UpdateStackRequestBody
	StackID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateStackRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, r.url(), body)
}

func (r *UpdateStackRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateStackRequest) url() string {
	return fmt.Sprintf("/v2/stacks/%s", url.PathEscape(r.StackID))
}

func (r *UpdateStackRequest) query() url.Values {
	return nil
}
