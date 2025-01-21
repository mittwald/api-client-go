package commonsv1

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "type":
//        type: "string"
//        enum:
//            - "ValidationError"
//    "message":
//        type: "string"
//        example: "Validation failed"
//    "validationErrors":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.commons.ValidationErrorSchema"}
//required:
//    - "type"
//    - "validationErrors"

type ValidationErrors struct {
	Type             ValidationErrorsType    `json:"type"`
	Message          *string                 `json:"message,omitempty"`
	ValidationErrors []ValidationErrorSchema `json:"validationErrors"`
}

func (o *ValidationErrors) Validate() error {
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	if o.ValidationErrors == nil {
		return errors.New("property validationErrors is required, but not set")
	}
	if err := func() error {
		for i := range o.ValidationErrors {
			if err := o.ValidationErrors[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property validationErrors: %w", err)
	}
	return nil
}
