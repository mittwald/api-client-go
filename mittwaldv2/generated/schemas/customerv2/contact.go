package customerv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "address": {"$ref": "#/components/schemas/de.mittwald.v1.commons.Address"}
//    "company":
//        type: "string"
//    "emailAddress":
//        type: "string"
//        format: "email"
//    "firstName":
//        type: "string"
//    "lastName":
//        type: "string"
//    "phoneNumbers":
//        type: "array"
//        items:
//            type: "string"
//    "salutation": {"$ref": "#/components/schemas/de.mittwald.v1.commons.Salutation"}
//    "title":
//        type: "string"
//    "useFormalTerm":
//        type: "boolean"
// required:
//    - "address"
//    - "salutation"

type Contact struct {
	Address       commonsv2.Address    `json:"address"`
	Company       *string              `json:"company,omitempty"`
	EmailAddress  *string              `json:"emailAddress,omitempty"`
	FirstName     *string              `json:"firstName,omitempty"`
	LastName      *string              `json:"lastName,omitempty"`
	PhoneNumbers  []string             `json:"phoneNumbers,omitempty"`
	Salutation    commonsv2.Salutation `json:"salutation"`
	Title         *string              `json:"title,omitempty"`
	UseFormalTerm *bool                `json:"useFormalTerm,omitempty"`
}

func (o *Contact) Validate() error {
	if err := o.Address.Validate(); err != nil {
		return fmt.Errorf("invalid property address: %w", err)
	}
	if err := func() error {
		if o.PhoneNumbers == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property phoneNumbers: %w", err)
	}
	if err := o.Salutation.Validate(); err != nil {
		return fmt.Errorf("invalid property salutation: %w", err)
	}
	return nil
}
