package marketplace

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev1"
)

type ListExtensionInstancesRequest struct {
	Limit     *int64
	Skip      *int64
	Page      *int64
	Context   marketplacev1.Context
	ContextID uuid.UUID
}

func (r *ListExtensionInstancesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListExtensionInstancesRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListExtensionInstancesRequest) url() string {
	return "/v2/extension-instances"
}

func (r *ListExtensionInstancesRequest) query() url.Values {
	q := make(url.Values)
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	q.Set("context", string(r.Context))
	q.Set("contextId", r.ContextID.String())
	return q
}
