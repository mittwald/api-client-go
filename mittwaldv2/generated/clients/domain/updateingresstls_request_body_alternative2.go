package domain

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

type UpdateIngressTLSRequestBodyAlternative2 struct {
	CertificateId uuid.UUID `json:"certificateId"`
}

func (o *UpdateIngressTLSRequestBodyAlternative2) Validate() error {
	return nil
}
