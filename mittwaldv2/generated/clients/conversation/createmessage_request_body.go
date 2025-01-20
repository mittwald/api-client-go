package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//properties:
//    "fileIds":
//        type: "array"
//        items:
//            type: "string"
//            format: "uuid"
//    "messageContent":
//        type: "string"
//        maxLength: 8000

//
type CreateMessageRequestBody struct {
	FileIds        []uuid.UUID `json:"fileIds,omitempty"`
	MessageContent *string     `json:"messageContent,omitempty"`
}

func (o *CreateMessageRequestBody) Validate() error {
	if err := func() error {
		if o.FileIds == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property fileIds: %w", err)
	}
	return nil
}