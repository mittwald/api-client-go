package conversation

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListConversationsRequest models a request for the
// 'conversation-list-conversations' operation. See [1] for more information.
//
// Get all conversation the authenticated user has created or has access to.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-list-conversations
type ListConversationsRequest struct {
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListConversationsRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *ListConversationsRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *ListConversationsRequest) url() string {
	return "/v2/conversations"
}

func (r *ListConversationsRequest) query() url.Values {
	return nil
}
