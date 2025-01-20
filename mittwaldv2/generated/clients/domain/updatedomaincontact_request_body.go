package domain

import (
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/domainv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "contact":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleField"}
//        minItems: 1
//required:
//    - "contact"

//
type UpdateDomainContactRequestBody struct {
	Contact []domainv1.HandleField `json:"contact"`
}

func (o *UpdateDomainContactRequestBody) Validate() error {
	if o.Contact == nil {
		return errors.New("property contact is required, but not set")
	}
	if err := func() error {
		for i := range o.Contact {
			if err := o.Contact[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property contact: %w", err)
	}
	return nil
}
