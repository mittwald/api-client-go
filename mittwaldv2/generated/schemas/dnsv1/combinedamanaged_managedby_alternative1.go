package dnsv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "ingressId":
//        type: "string"
//        format: "uuid"
//required:
//    - "ingressId"

//
type CombinedAManagedManagedByAlternative1 struct {
	IngressId uuid.UUID `json:"ingressId"`
}

func (o *CombinedAManagedManagedByAlternative1) Validate() error {
	return nil
}