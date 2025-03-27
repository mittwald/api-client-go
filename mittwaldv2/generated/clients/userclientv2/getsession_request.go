package userclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetSessionRequest models a request for the 'user-get-session' operation. See [1]
// for more information.
//
// Get a specific session.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-get-session
type GetSessionRequest struct {
	TokenID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetSessionRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
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

func (r *GetSessionRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetSessionRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/users/self/sessions/%s", url.PathEscape(r.TokenID)),
	}
	return u.String()
}

func (r *GetSessionRequest) query() url.Values {
	return nil
}
