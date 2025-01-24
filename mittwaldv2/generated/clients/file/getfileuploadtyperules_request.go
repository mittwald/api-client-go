package file

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type GetFileUploadTypeRulesRequest struct {
	FileUploadType GetFileUploadTypeRulesRequestPathFileUploadType
}

func (r *GetFileUploadTypeRulesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetFileUploadTypeRulesRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetFileUploadTypeRulesRequest) url() string {
	return fmt.Sprintf("/v2/file-upload-types/%s/rules", url.PathEscape(string(r.FileUploadType)))
}

func (r *GetFileUploadTypeRulesRequest) query() url.Values {
	return nil
}
