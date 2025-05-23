package sslv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "commonName":
//        type: "string"
//    "contact": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.Contact"}
//    "projectId":
//        type: "string"
//        format: "uuid"
// required:
//    - "projectId"
//    - "commonName"
//    - "contact"

type CertificateRequestCreateWithCSRRequest struct {
	CommonName string  `json:"commonName"`
	Contact    Contact `json:"contact"`
	ProjectId  string  `json:"projectId"`
}

func (o *CertificateRequestCreateWithCSRRequest) Validate() error {
	if err := o.Contact.Validate(); err != nil {
		return fmt.Errorf("invalid property contact: %w", err)
	}
	return nil
}
