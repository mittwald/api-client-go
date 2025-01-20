package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/ingressv1"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "ownership": {"$ref": "#/components/schemas/de.mittwald.v1.ingress.Ownership"}
//required:
//    - "id"
//    - "ownership"

//
type CreateIngressResponse struct {
	Id        uuid.UUID           `json:"id"`
	Ownership ingressv1.Ownership `json:"ownership"`
}

func (o *CreateIngressResponse) Validate() error {
	if err := o.Ownership.Validate(); err != nil {
		return fmt.Errorf("invalid property ownership: %w", err)
	}
	return nil
}
