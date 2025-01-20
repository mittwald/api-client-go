package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type UpdateConversationRequest struct {
	Body           UpdateConversationRequestBody
	ConversationID uuid.UUID
}

func (r *UpdateConversationRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *UpdateConversationRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *UpdateConversationRequest) url() string {
	return fmt.Sprintf("/v2/conversations/%s", url.PathEscape(r.ConversationID.String()))
}

func (r *UpdateConversationRequest) query() url.Values {
	return nil
}