package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListAppsRequest models a request for the 'app-list-apps' operation. See [1] for
// more information.
//
// List Apps.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/app/app-list-apps
type ListAppsRequest struct {
	Limit *int64
	Skip  *int64
	Page  *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListAppsRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ListAppsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListAppsRequest) url() string {
	u := url.URL{
		Path:     "/v2/apps",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListAppsRequest) query() url.Values {
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
	return q
}
