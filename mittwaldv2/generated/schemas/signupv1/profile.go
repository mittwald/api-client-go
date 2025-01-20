package signupv1

import (
	"fmt"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "email":
//        type: "string"
//        format: "email"
//        example: "a.lovelace@example.com"
//    "mfaDetails":
//        type: "object"
//        properties:
//            "mfaConfirmed":
//                type: "boolean"
//            "mfaInitialized":
//                type: "boolean"
//        description: "the users mfa details"
//    "passwordUpdatedAt":
//        type: "string"
//        format: "date-time"
//    "person": {"$ref": "#/components/schemas/de.mittwald.v1.commons.Person"}
//    "userId":
//        type: "string"

//
type Profile struct {
	Email             *string            `json:"email,omitempty"`
	MfaDetails        *ProfileMFADetails `json:"mfaDetails,omitempty"`
	PasswordUpdatedAt *time.Time         `json:"passwordUpdatedAt,omitempty"`
	Person            *commonsv1.Person  `json:"person,omitempty"`
	UserId            *string            `json:"userId,omitempty"`
}

func (o *Profile) Validate() error {
	if err := func() error {
		if o.MfaDetails == nil {
			return nil
		}
		return o.MfaDetails.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property mfaDetails: %w", err)
	}
	if err := func() error {
		if o.Person == nil {
			return nil
		}
		return o.Person.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property person: %w", err)
	}
	return nil
}
