package user

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
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/deprecated-user-service-issue-new
type DeprecatedServiceIssueNewRequest struct {
	Body DeprecatedServiceIssueNewRequestBody
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DeprecatedServiceIssueNewRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *DeprecatedServiceIssueNewRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *DeprecatedServiceIssueNewRequest) url() string {
	return "/v2/user/issues"
}

func (r *DeprecatedServiceIssueNewRequest) query() url.Values {
	return nil
}
