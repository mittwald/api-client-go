package projectclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DeleteProjectRequest models a request for the 'project-delete-project'
// operation. See [1] for more information.
//
// Delete a Project.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-delete-project
type DeleteProjectRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeleteProjectRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, r.url(), body)
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

func (r *DeleteProjectRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DeleteProjectRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *DeleteProjectRequest) query() url.Values {
	return nil
}
