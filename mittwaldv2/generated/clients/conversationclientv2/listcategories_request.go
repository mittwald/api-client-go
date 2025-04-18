package conversationclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListCategoriesRequest models a request for the 'conversation-list-categories'
// operation. See [1] for more information.
//
// Get all conversation categories.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-list-categories
type ListCategoriesRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListCategoriesRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListCategoriesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListCategoriesRequest) url() string {
	u := url.URL{
		Path: "/v2/conversation-categories",
	}
	return u.String()
}

func (r *ListCategoriesRequest) query() url.Values {
	return nil
}
