package domainv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "authCode": {"$ref": "#/components/schemas/de.mittwald.v1.domain.AuthCode"}
//    "authCode2": {"$ref": "#/components/schemas/de.mittwald.v1.domain.AuthCode2"}
//    "connected":
//        type: "boolean"
//    "contactHash":
//        type: "string"
//    "deleted":
//        type: "boolean"
//    "domain":
//        type: "string"
//        format: "naked-domain"
//    "domainId":
//        type: "string"
//        format: "uuid"
//    "handles":
//        type: "object"
//        properties:
//            "adminC": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleReadable"}
//            "ownerC": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleReadable"}
//        required:
//            - "ownerC"
//    "nameservers":
//        type: "array"
//        items:
//            type: "string"
//            format: "idn-hostname"
//        uniqueItems: true
//    "processes":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.domain.Process"}
//    "projectId":
//        type: "string"
//        format: "uuid"
//    "transferInAuthCode":
//        type: "string"
//    "usesDefaultNameserver":
//        type: "boolean"
//required:
//    - "domainId"
//    - "domain"
//    - "projectId"
//    - "deleted"
//    - "nameservers"
//    - "handles"
//    - "connected"
//    - "usesDefaultNameserver"
//    - "hasAuthCode"

//
type Domain struct {
	AuthCode              *AuthCode     `json:"authCode,omitempty"`
	AuthCode2             *AuthCode2    `json:"authCode2,omitempty"`
	Connected             bool          `json:"connected"`
	ContactHash           *string       `json:"contactHash,omitempty"`
	Deleted               bool          `json:"deleted"`
	Domain                string        `json:"domain"`
	DomainId              uuid.UUID     `json:"domainId"`
	Handles               DomainHandles `json:"handles"`
	Nameservers           []string      `json:"nameservers"`
	Processes             []Process     `json:"processes,omitempty"`
	ProjectId             uuid.UUID     `json:"projectId"`
	TransferInAuthCode    *string       `json:"transferInAuthCode,omitempty"`
	UsesDefaultNameserver bool          `json:"usesDefaultNameserver"`
}

func (o *Domain) Validate() error {
	if err := func() error {
		if o.AuthCode == nil {
			return nil
		}
		return o.AuthCode.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property authCode: %w", err)
	}
	if err := func() error {
		if o.AuthCode2 == nil {
			return nil
		}
		return o.AuthCode2.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property authCode2: %w", err)
	}
	if err := o.Handles.Validate(); err != nil {
		return fmt.Errorf("invalid property handles: %w", err)
	}
	if o.Nameservers == nil {
		return errors.New("property nameservers is required, but not set")
	}
	if err := func() error {
		if o.Processes == nil {
			return nil
		}
		return func() error {
			for i := range o.Processes {
				if err := o.Processes[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property processes: %w", err)
	}
	return nil
}
