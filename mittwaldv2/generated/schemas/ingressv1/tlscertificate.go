package ingressv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "certificateId":
//        type: "string"
//        format: "uuid"
//required:
//    - "certificateId"
//additionalProperties: false

//
type TlsCertificate struct {
	CertificateId uuid.UUID `json:"certificateId"`
}

func (o *TlsCertificate) Validate() error {
	return nil
}