package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type GetConversationRequest struct {
	ConversationID uuid.UUID
}

func (r *GetConversationRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetConversationRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetConversationRequest) url() string {
	return fmt.Sprintf("/v2/conversations/%s", url.PathEscape(r.ConversationID.String()))
}

func (r *GetConversationRequest) query() url.Values {
	return nil
}
