package userclientv2

import "errors"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "recoveryCodesList":
//        type: "array"
//        items:
//            type: "string"
//            maxLength: 16
//            minLength: 16
//            example: 1234123412341234
//        maxItems: 20
//        minItems: 20
// required:
//    - "recoveryCodesList"

type ConfirmMFAResponse struct {
	RecoveryCodesList []string `json:"recoveryCodesList"`
}

func (o *ConfirmMFAResponse) Validate() error {
	if o.RecoveryCodesList == nil {
		return errors.New("property recoveryCodesList is required, but not set")
	}
	return nil
}
