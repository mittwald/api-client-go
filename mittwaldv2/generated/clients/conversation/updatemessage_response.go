package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "messageId":
//        type: "string"
//        format: "uuid"
//required:
//    - "messageId"

//
type UpdateMessageResponse struct {
	MessageId uuid.UUID `json:"messageId"`
}

func (o *UpdateMessageResponse) Validate() error {
	return nil
}
