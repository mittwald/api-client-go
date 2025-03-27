package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetProjectRequest models a request for the 'project-get-project' operation. See
// [1] for more information.
//
// Get a Project.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/project/project-get-project
type GetProjectRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetProjectRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetProjectRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetProjectRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *GetProjectRequest) query() url.Values {
	return nil
}
