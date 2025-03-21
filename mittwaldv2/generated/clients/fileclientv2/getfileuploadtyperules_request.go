package fileclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetFileUploadTypeRulesRequest models a request for the
// 'file-get-file-upload-type-rules' operation. See [1] for more information.
//
// Get a FileUploadType's rules.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/file/file-get-file-upload-type-rules
type GetFileUploadTypeRulesRequest struct {
	FileUploadType GetFileUploadTypeRulesRequestPathFileUploadType
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetFileUploadTypeRulesRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetFileUploadTypeRulesRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetFileUploadTypeRulesRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/file-upload-types/%s/rules", url.PathEscape(string(r.FileUploadType))),
	}
	return u.String()
}

func (r *GetFileUploadTypeRulesRequest) query() url.Values {
	return nil
}
