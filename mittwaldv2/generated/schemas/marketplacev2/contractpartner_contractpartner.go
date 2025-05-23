package marketplacev2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "address": {"$ref": "#/components/schemas/de.mittwald.v1.commons.Address"}
//    "company":
//        type: "string"
//        example: "mittwald"
//    "email":
//        type: "string"
//        example: "a.lovelace@example.com"
//    "firstName":
//        type: "string"
//        example: "Ada"
//    "lastName":
//        type: "string"
//        example: "Lovelace"
// required:
//    - "firstName"
//    - "lastName"
//    - "email"
//    - "address"

type ContractPartnerContractPartner struct {
	Address   commonsv2.Address `json:"address"`
	Company   *string           `json:"company,omitempty"`
	Email     string            `json:"email"`
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
}

func (o *ContractPartnerContractPartner) Validate() error {
	if err := o.Address.Validate(); err != nil {
		return fmt.Errorf("invalid property address: %w", err)
	}
	return nil
}
