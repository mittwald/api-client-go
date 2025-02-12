package conversationclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// GetConversationMembersRequest models a request for the
// 'conversation-get-conversation-members' operation. See [1] for more information.
//
// Get members of a support conversation.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-get-conversation-members
type GetConversationMembersRequest struct {
	ConversationID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *GetConversationMembersRequest) BuildRequest() (*http.Request, error) {
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

func (r *GetConversationMembersRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *GetConversationMembersRequest) url() string {
	return fmt.Sprintf("/v2/conversations/%s/members", url.PathEscape(r.ConversationID))
}

func (r *GetConversationMembersRequest) query() url.Values {
	return nil
}
