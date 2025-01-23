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

type DeprecatedUpdateProjectMailSettingRequest struct {
	Body      DeprecatedUpdateProjectMailSettingRequestBody
	ProjectID string
	Setting   DeprecatedUpdateProjectMailSettingRequestPathSetting
}

func (r *DeprecatedUpdateProjectMailSettingRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *DeprecatedUpdateProjectMailSettingRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedUpdateProjectMailSettingRequest) url() string {
	return fmt.Sprintf("/v2/projects/%s/mail-settings/%s", url.PathEscape(r.ProjectID), url.PathEscape(string(r.Setting)))
}

func (r *DeprecatedUpdateProjectMailSettingRequest) query() url.Values {
	return nil
}