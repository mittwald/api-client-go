package cronjobv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "id":
//        type: "string"
//        format: "uuid"

type CronjobExecutionAbortedBy struct {
	Id *uuid.UUID `json:"id,omitempty"`
}

func (o *CronjobExecutionAbortedBy) Validate() error {
	return nil
}
