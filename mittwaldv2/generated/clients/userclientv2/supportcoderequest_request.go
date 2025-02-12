package userclientv2

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// SupportCodeRequestRequest models a request for the 'user-support-code-request'
// operation. See [1] for more information.
//
// Request a support code.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/user/user-support-code-request
type SupportCodeRequestRequest struct {
	ForceRecreate *bool
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *SupportCodeRequestRequest) BuildRequest() (*http.Request, error) {
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

func (r *SupportCodeRequestRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *SupportCodeRequestRequest) url() string {
	return "/v2/users/self/credentials/support-code"
}

func (r *SupportCodeRequestRequest) query() url.Values {
	q := make(url.Values)
	if r.ForceRecreate != nil {
		q.Set("forceRecreate", strconv.FormatBool(*r.ForceRecreate))
	}
	return q
}
