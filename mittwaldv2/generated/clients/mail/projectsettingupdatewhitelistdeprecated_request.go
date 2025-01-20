package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type ProjectsettingUpdateWhitelistDeprecatedRequest struct {
	Body      ProjectsettingUpdateWhitelistDeprecatedRequestBody
	ProjectID string
}

func (r *ProjectsettingUpdateWhitelistDeprecatedRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *ProjectsettingUpdateWhitelistDeprecatedRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *ProjectsettingUpdateWhitelistDeprecatedRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/mailsettings/whitelist", url.PathEscape(r.ProjectID))
}

func (r *ProjectsettingUpdateWhitelistDeprecatedRequest) query() url.Values {
	return nil
}