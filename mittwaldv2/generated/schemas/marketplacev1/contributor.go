package marketplacev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "customerId":
//        type: "string"
//        format: "uuid"
//    "description":
//        type: "string"
//    "email":
//        type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "logoRefId":
//        type: "string"
//    "name":
//        type: "string"
//    "phone":
//        type: "string"
//    "state": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ContributorState"}
//    "url":
//        type: "string"
//required:
//    - "id"
//    - "customerId"
//    - "state"
//    - "name"
//    - "description"

//
type Contributor struct {
	CustomerId  uuid.UUID        `json:"customerId"`
	Description string           `json:"description"`
	Email       *string          `json:"email,omitempty"`
	Id          uuid.UUID        `json:"id"`
	LogoRefId   *string          `json:"logoRefId,omitempty"`
	Name        string           `json:"name"`
	Phone       *string          `json:"phone,omitempty"`
	State       ContributorState `json:"state"`
	Url         *string          `json:"url,omitempty"`
}

func (o *Contributor) Validate() error {
	if err := o.State.Validate(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	return nil
}
