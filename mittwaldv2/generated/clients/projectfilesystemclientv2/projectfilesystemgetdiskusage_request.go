package projectfilesystemclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ProjectFileSystemGetDiskUsageRequest models a request for the
// 'project-file-system-get-disk-usage' operation. See [1] for more information.
//
// Get a Project directory filesystem usage.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/project file
// system/project-file-system-get-disk-usage
type ProjectFileSystemGetDiskUsageRequest struct {
	ProjectID string
	Directory *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ProjectFileSystemGetDiskUsageRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *ProjectFileSystemGetDiskUsageRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ProjectFileSystemGetDiskUsageRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/filesystem-disk-usage", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ProjectFileSystemGetDiskUsageRequest) query() url.Values {
	q := make(url.Values)
	if r.Directory != nil {
		q.Set("directory", *r.Directory)
	}
	return q
}
