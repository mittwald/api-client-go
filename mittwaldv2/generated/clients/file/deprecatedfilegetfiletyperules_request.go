package file

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type DeprecatedFileGetFileTypeRulesRequest struct {
	Name DeprecatedFileGetFileTypeRulesRequestPathName
}

func (r *DeprecatedFileGetFileTypeRulesRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *DeprecatedFileGetFileTypeRulesRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *DeprecatedFileGetFileTypeRulesRequest) url() string {
	return fmt.Sprintf("/v2/file-type-rules/%s", url.PathEscape(string(r.Name)))
}

func (r *DeprecatedFileGetFileTypeRulesRequest) query() url.Values {
	return nil
}
