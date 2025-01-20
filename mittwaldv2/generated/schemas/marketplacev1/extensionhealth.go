package marketplacev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "extensionInstances":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionInstanceHealth"}
//    "functional":
//        type: "boolean"
//        default: false
//    "id":
//        type: "string"
//        format: "uuid"
//    "inoperableReason":
//        type: "string"
//        example: "9 of 10 webhooks in the last hour were unsuccessful."
//    "published":
//        type: "boolean"
//        default: false
//    "withdrawalReason":
//        type: "string"
//required:
//    - "id"
//    - "functional"
//    - "published"
//    - "extensionInstances"

//
type ExtensionHealth struct {
	ExtensionInstances []ExtensionInstanceHealth `json:"extensionInstances"`
	Functional         bool                      `json:"functional"`
	Id                 uuid.UUID                 `json:"id"`
	InoperableReason   *string                   `json:"inoperableReason,omitempty"`
	Published          bool                      `json:"published"`
	WithdrawalReason   *string                   `json:"withdrawalReason,omitempty"`
}

func (o *ExtensionHealth) Validate() error {
	if o.ExtensionInstances == nil {
		return errors.New("property extensionInstances is required, but not set")
	}
	if err := func() error {
		for i := range o.ExtensionInstances {
			if err := o.ExtensionInstances[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property extensionInstances: %w", err)
	}
	return nil
}