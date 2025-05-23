package sslv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "caBundle":
//        type: "string"
//    "certificate":
//        type: "string"
//    "certificateOrderId":
//        type: "string"
//        format: "uuid"
//    "certificateRequestId":
//        type: "string"
//        format: "uuid"
//    "certificateType": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateType"}
//    "commonName":
//        type: "string"
//    "contact": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.Contact"}
//    "dnsNames":
//        type: "array"
//        items:
//            type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "isExpired":
//        type: "boolean"
//    "issuer":
//        type: "string"
//    "lastExpirationThresholdHit":
//        type: "integer"
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
//    - "certificateRequestId"
//    - "projectId"
//    - "validFrom"
//    - "validTo"
//    - "certificate"
//    - "isExpired"
//    - "lastExpirationThresholdHit"

type Certificate struct {
	CaBundle                   *string   `json:"caBundle,omitempty"`
	Certificate                string    `json:"certificate"`
	CertificateOrderId         *string   `json:"certificateOrderId,omitempty"`
	CertificateRequestId       string    `json:"certificateRequestId"`
	CertificateType            int64     `json:"certificateType"`
	CommonName                 *string   `json:"commonName,omitempty"`
	Contact                    *Contact  `json:"contact,omitempty"`
	DnsNames                   []string  `json:"dnsNames,omitempty"`
	Id                         string    `json:"id"`
	IsExpired                  bool      `json:"isExpired"`
	Issuer                     *string   `json:"issuer,omitempty"`
	LastExpirationThresholdHit int64     `json:"lastExpirationThresholdHit"`
	ProjectId                  string    `json:"projectId"`
	ValidFrom                  time.Time `json:"validFrom"`
	ValidTo                    time.Time `json:"validTo"`
}

func (o *Certificate) Validate() error {
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
