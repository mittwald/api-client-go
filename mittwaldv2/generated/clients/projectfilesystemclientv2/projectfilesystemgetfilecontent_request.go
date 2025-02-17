package projectfilesystemclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ProjectFileSystemGetFileContentRequest models a request for the
// 'project-file-system-get-file-content' operation. See [1] for more information.
//
// Get a Project file's content.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/project file
// system/project-file-system-get-file-content
type ProjectFileSystemGetFileContentRequest struct {
	ProjectID string
	File      *string
	Inline    *bool
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ProjectFileSystemGetFileContentRequest) BuildRequest() (*http.Request, error) {
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

func (r *ProjectFileSystemGetFileContentRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ProjectFileSystemGetFileContentRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/projects/%s/filesystem-file-content", url.PathEscape(r.ProjectID)),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ProjectFileSystemGetFileContentRequest) query() url.Values {
	q := make(url.Values)
	if r.File != nil {
		q.Set("file", *r.File)
	}
	if r.Inline != nil {
		q.Set("inline", strconv.FormatBool(*r.Inline))
	}
	return q
}
