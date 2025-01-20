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

type SetConversationStatusRequest struct {
	Body           SetConversationStatusRequestBody
	ConversationID uuid.UUID
}

func (r *SetConversationStatusRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, r.url(), body)
}

func (r *SetConversationStatusRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *SetConversationStatusRequest) url() string {
	return fmt.Sprintf("/v2/conversations/%s/status", url.PathEscape(r.ConversationID.String()))
}

func (r *SetConversationStatusRequest) query() url.Values {
	return nil
}
