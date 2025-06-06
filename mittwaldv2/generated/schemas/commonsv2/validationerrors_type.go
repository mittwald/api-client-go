package commonsv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "ValidationError"

type ValidationErrorsType string

const ValidationErrorsTypeValidationError ValidationErrorsType = "ValidationError"

func (e ValidationErrorsType) Validate() error {
	if e == ValidationErrorsTypeValidationError {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
