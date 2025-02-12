package projectfilesystemclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ProjectFileSystemGetJwtRequest models a request for the
// 'project-file-system-get-jwt' operation. See [1] for more information.
//
// Get a Project's file/filesystem authorization token.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/project file
// system/project-file-system-get-jwt
type ProjectFileSystemGetJwtRequest struct {
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ProjectFileSystemGetJwtRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *ProjectFileSystemGetJwtRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ProjectFileSystemGetJwtRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/jwt", url.PathEscape(r.ProjectID))
}

func (r *ProjectFileSystemGetJwtRequest) query() url.Values {
	return nil
}
