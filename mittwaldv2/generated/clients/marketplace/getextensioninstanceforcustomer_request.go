package marketplace

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetExtensionInstanceForCustomerRequest struct {
	CustomerID  string
	ExtensionID uuid.UUID
}

func (r *GetExtensionInstanceForCustomerRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetExtensionInstanceForCustomerRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetExtensionInstanceForCustomerRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/extensions/%s", url.PathEscape(r.CustomerID), url.PathEscape(r.ExtensionID.String()))
}

func (r *GetExtensionInstanceForCustomerRequest) query() url.Values {
	return nil
}
