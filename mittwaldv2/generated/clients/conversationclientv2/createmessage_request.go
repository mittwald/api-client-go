package conversationclientv2

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

// CreateMessageRequest models a request for the 'conversation-create-message'
// operation. See [1] for more information.
//
// Send a new message in the conversation.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-create-message
type CreateMessageRequest struct {
	Body           CreateMessageRequestBody
	ConversationID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *CreateMessageRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *CreateMessageRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *CreateMessageRequest) url() string {
	return fmt.Sprintf("/v2/conversations/%s/messages", url.PathEscape(r.ConversationID))
}

func (r *CreateMessageRequest) query() url.Values {
	return nil
}
