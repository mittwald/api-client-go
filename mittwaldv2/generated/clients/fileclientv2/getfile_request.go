package fileclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetFileRequest models a request for the 'file-get-file' operation. See [1] for
// more information.
//
// Get a File.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/file/file-get-file
type GetFileRequest struct {
	FileID             string
	Accept             *GetFileRequestQueryAccept
	ContentDisposition *GetFileRequestQueryContentDisposition
	Token              *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetFileRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetFileRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetFileRequest) url() string {
	return fmt.Sprintf("/v2/files/%s", url.PathEscape(r.FileID))
}

func (r *GetFileRequest) query() url.Values {
	q := make(url.Values)
	if r.Accept != nil {
		q.Set("accept", string(*r.Accept))
	}
	if r.ContentDisposition != nil {
		q.Set("content-disposition", string(*r.ContentDisposition))
	}
	if r.Token != nil {
		q.Set("token", *r.Token)
	}
	return q
}
