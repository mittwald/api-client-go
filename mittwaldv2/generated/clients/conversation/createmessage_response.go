package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "conversationId":
//        type: "string"
//        format: "uuid"
//    "messageId":
//        type: "string"
//        format: "uuid"
//required:
//    - "conversationId"
//    - "messageId"

type CreateMessageResponse struct {
	ConversationId uuid.UUID `json:"conversationId"`
	MessageId      uuid.UUID `json:"messageId"`
}

func (o *CreateMessageResponse) Validate() error {
	return nil
}
