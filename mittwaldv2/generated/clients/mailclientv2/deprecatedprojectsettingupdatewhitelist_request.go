package mailclientv2

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

// DeprecatedProjectsettingUpdateWhitelistRequest models a request for the
// 'deprecated-mail-projectsetting-update-whitelist' operation. See [1] for more
// information.
//
// # Update whitelist for a given project ID
//
// This operation is deprecated. Use the PATCH
// v2/{projectId}/mail-settings/{mailSetting} endpoint instead.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/mail/deprecated-mail-projectsetting-update-whitelist
type DeprecatedProjectsettingUpdateWhitelistRequest struct {
	Body      DeprecatedProjectsettingUpdateWhitelistRequestBody
	ProjectID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedProjectsettingUpdateWhitelistRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, r.url(), body)
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

func (r *DeprecatedProjectsettingUpdateWhitelistRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedProjectsettingUpdateWhitelistRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/projects/%s/mailsettings/whitelist", url.PathEscape(r.ProjectID)),
	}
	return u.String()
}

func (r *DeprecatedProjectsettingUpdateWhitelistRequest) query() url.Values {
	return nil
}
