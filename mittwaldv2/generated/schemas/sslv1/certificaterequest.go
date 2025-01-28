package sslv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "certificateData": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateData"}
//    "certificateType": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateType"}
//    "commonName":
//        type: "string"
//    "contact": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.Contact"}
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "dnsNames":
//        type: "array"
//        items:
//            type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "isCompleted":
//        type: "boolean"
//    "issuer":
//        type: "string"
//    "projectId":
//        type: "string"
//        format: "uuid"
//    "validFrom":
//        type: "string"
//        format: "date-time"
//    "validTo":
//        type: "string"
//        format: "date-time"
// required:
//    - "id"
//    - "certificateType"
//    - "projectId"
//    - "certificateData"
//    - "isCompleted"
//    - "createdAt"

type CertificateRequest struct {
	CertificateData CertificateData `json:"certificateData"`
	CertificateType int64           `json:"certificateType"`
	CommonName      *string         `json:"commonName,omitempty"`
	Contact         *Contact        `json:"contact,omitempty"`
	CreatedAt       time.Time       `json:"createdAt"`
	DnsNames        []string        `json:"dnsNames,omitempty"`
	Id              string          `json:"id"`
	IsCompleted     bool            `json:"isCompleted"`
	Issuer          *string         `json:"issuer,omitempty"`
	ProjectId       string          `json:"projectId"`
	ValidFrom       *time.Time      `json:"validFrom,omitempty"`
	ValidTo         *time.Time      `json:"validTo,omitempty"`
}

func (o *CertificateRequest) Validate() error {
	if err := o.CertificateData.Validate(); err != nil {
		return fmt.Errorf("invalid property certificateData: %w", err)
	}
	if err := func() error {
		if o.Contact == nil {
			return nil
		}
		return o.Contact.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property contact: %w", err)
	}
	if err := func() error {
		if o.DnsNames == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property dnsNames: %w", err)
	}
	return nil
}
