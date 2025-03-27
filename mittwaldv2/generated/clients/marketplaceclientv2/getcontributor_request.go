package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetContributorRequest models a request for the 'extension-get-contributor'
// operation. See [1] for more information.
//
// Get a Contributor.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-get-contributor
type GetContributorRequest struct {
	ContributorID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetContributorRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetContributorRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetContributorRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/contributors/%s", url.PathEscape(r.ContributorID)),
	}
	return u.String()
}

func (r *GetContributorRequest) query() url.Values {
	return nil
}
