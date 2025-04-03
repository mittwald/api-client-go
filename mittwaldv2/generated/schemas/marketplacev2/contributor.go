package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "customerId":
//        type: "string"
//        format: "uuid"
//    "description":
//        type: "string"
//    "email":
//        type: "string"
//        deprecated: true
//    "id":
//        type: "string"
//        format: "uuid"
//    "logoRefId":
//        type: "string"
//    "name":
//        type: "string"
//    "phone":
//        type: "string"
//        deprecated: true
//    "state": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ContributorState"}
//    "supportInformation": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SupportMeta"}
//    "url":
//        type: "string"
// required:
//    - "id"
//    - "customerId"
//    - "state"
//    - "name"
//    - "supportInformation"

type Contributor struct {
	CustomerId         string           `json:"customerId"`
	Description        *string          `json:"description,omitempty"`
	Email              *string          `json:"email,omitempty"`
	Id                 string           `json:"id"`
	LogoRefId          *string          `json:"logoRefId,omitempty"`
	Name               string           `json:"name"`
	Phone              *string          `json:"phone,omitempty"`
	State              ContributorState `json:"state"`
	SupportInformation SupportMeta      `json:"supportInformation"`
	Url                *string          `json:"url,omitempty"`
}

func (o *Contributor) Validate() error {
	if err := o.State.Validate(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	if err := o.SupportInformation.Validate(); err != nil {
		return fmt.Errorf("invalid property supportInformation: %w", err)
	}
	return nil
}
