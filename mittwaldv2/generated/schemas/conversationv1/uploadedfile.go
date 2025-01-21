package conversationv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "name":
//        type: "string"
//    "status":
//        type: "string"
//        enum:
//            - "uploaded"
//    "type":
//        type: "string"
//required:
//    - "id"
//    - "status"
//    - "name"
//    - "type"

type UploadedFile struct {
	Id     uuid.UUID          `json:"id"`
	Name   string             `json:"name"`
	Status UploadedFileStatus `json:"status"`
	Type   string             `json:"type"`
}

func (o *UploadedFile) Validate() error {
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	return nil
}
