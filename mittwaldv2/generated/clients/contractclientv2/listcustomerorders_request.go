package contractclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListCustomerOrdersRequest models a request for the 'order-list-customer-orders'
// operation. See [1] for more information.
//
// Get list of Orders of a Customer.
//
// The list of Orders of a Customer the User has access to, can be filtered by
// orderStatus, articleTemplate and count.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/contract/order-list-customer-orders
type ListCustomerOrdersRequest struct {
	CustomerID     string
	Limit          *int64
	Skip           *int64
	Page           *int64
	IncludesStatus []orderv2.OrderStatus
	ExcludesStatus []orderv2.OrderStatus
	TemplateNames  []string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListCustomerOrdersRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
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

func (r *ListCustomerOrdersRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListCustomerOrdersRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/customers/%s/orders", url.PathEscape(r.CustomerID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListCustomerOrdersRequest) query() url.Values {
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
