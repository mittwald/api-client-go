package userclientv2

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

// DeprecatedServiceIssueNewRequest models a request for the
// 'deprecated-user-service-issue-new' operation. See [1] for more information.
//
// create a new issue
//
// was replaced by /v2/users/self/feedback.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-issue-new
type DeprecatedServiceIssueNewRequest struct {
	Body DeprecatedServiceIssueNewRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServiceIssueNewRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
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

func (r *DeprecatedServiceIssueNewRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *DeprecatedServiceIssueNewRequest) url() string {
	u := url.URL{
		Path: "/v2/user/issues",
	}
	return u.String()
}

func (r *DeprecatedServiceIssueNewRequest) query() url.Values {
	return nil
}
