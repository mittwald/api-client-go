package app

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "description":
//        type: "string"
//    "targetProjectId":
//        type: "string"
//        format: "uuid"
//required:
//    - "description"

//
type RequestAppinstallationCopyRequestBody struct {
	Description     string     `json:"description"`
	TargetProjectId *uuid.UUID `json:"targetProjectId,omitempty"`
}

func (o *RequestAppinstallationCopyRequestBody) Validate() error {
	return nil
}
