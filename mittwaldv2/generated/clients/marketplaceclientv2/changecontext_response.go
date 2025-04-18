package marketplaceclientv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "cleaningUpInstances":
//        type: "boolean"
//        description: "If this value is true the context will change asynchronously after removing all extension-instances of this extension."
//        example: true
//    "currentContext": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.Context"}
//    "id":
//        type: "string"
//        format: "uuid"
// required:
//    - "id"
//    - "currentContext"
//    - "cleaningUpInstances"

type ChangeContextResponse struct {
	CleaningUpInstances bool                  `json:"cleaningUpInstances"`
	CurrentContext      marketplacev2.Context `json:"currentContext"`
	Id                  string                `json:"id"`
}

func (o *ChangeContextResponse) Validate() error {
	if err := o.CurrentContext.Validate(); err != nil {
		return fmt.Errorf("invalid property currentContext: %w", err)
	}
	return nil
}
