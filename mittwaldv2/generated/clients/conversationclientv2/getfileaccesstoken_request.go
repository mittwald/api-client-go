package conversationclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetFileAccessTokenRequest models a request for the
// 'conversation-get-file-access-token' operation. See [1] for more information.
//
// Request an access token for the File belonging to the Conversation.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-get-file-access-token
type GetFileAccessTokenRequest struct {
	ConversationID string
	FileID         string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetFileAccessTokenRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
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

func (r *GetFileAccessTokenRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetFileAccessTokenRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/conversations/%s/files/%s/access-token", url.PathEscape(r.ConversationID), url.PathEscape(r.FileID)),
	}
	return u.String()
}

func (r *GetFileAccessTokenRequest) query() url.Values {
	return nil
}
