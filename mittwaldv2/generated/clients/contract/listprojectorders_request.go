package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv1"
)

type ListProjectOrdersRequest struct {
	ProjectID      uuid.UUID
	Limit          *int64
	Skip           *int64
	Page           *int64
	IncludesStatus []orderv1.OrderStatus
	ExcludesStatus []orderv1.OrderStatus
	TemplateNames  []string
}

func (r *ListProjectOrdersRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListProjectOrdersRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListProjectOrdersRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/orders", url.PathEscape(r.ProjectID.String()))
}

func (r *ListProjectOrdersRequest) query() url.Values {
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
	for _, val := range r.IncludesStatus {
		q.Add("includesStatus", string(val))
	}
	for _, val := range r.ExcludesStatus {
		q.Add("excludesStatus", string(val))
	}
	for _, val := range r.TemplateNames {
		q.Add("templateNames", val)
	}
	return q
}
