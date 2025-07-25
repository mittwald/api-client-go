package leadfyndrv2

import (
	"errors"
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "actualUrl":
//        type: "string"
//    "businessFields":
//        type: "array"
//        items:
//            type: "string"
//    "company": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.DetailCompany"}
//    "contact": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Contact"}
//    "description":
//        type: "string"
//    "domain":
//        type: "string"
//    "hoster": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Hoster"}
//    "languages":
//        type: "array"
//        items:
//            type: "string"
//    "leadId":
//        type: "string"
//    "mainTechnology": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Technology"}
//    "metrics": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.DetailMetrics"}
//    "potential":
//        type: "number"
//        maximum: 1
//        minimum: 0
//        format: "float"
//    "reservationAllowed":
//        type: "boolean"
//    "reservedAt":
//        type: "string"
//        format: "date-time"
//    "scannedAt":
//        type: "string"
//        format: "date-time"
//    "screenshot":
//        type: "string"
//    "socialMedia":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.SocialMedia"}
//    "technologies":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Technology"}
//    "unlockedAt":
//        type: "string"
//        format: "date-time"
// required:
//    - "leadId"
//    - "potential"
//    - "screenshot"
//    - "company"
//    - "metrics"
//    - "businessFields"
//    - "description"
//    - "technologies"
//    - "domain"
//    - "actualUrl"
//    - "socialMedia"
//    - "contact"
//    - "hoster"
//    - "unlockedAt"
//    - "languages"

type UnlockedLead struct {
	ActualUrl          string        `json:"actualUrl"`
	BusinessFields     []string      `json:"businessFields"`
	Company            any           `json:"company"`
	Contact            Contact       `json:"contact"`
	Description        string        `json:"description"`
	Domain             string        `json:"domain"`
	Hoster             Hoster        `json:"hoster"`
	Languages          []string      `json:"languages"`
	LeadId             string        `json:"leadId"`
	MainTechnology     *Technology   `json:"mainTechnology,omitempty"`
	Metrics            DetailMetrics `json:"metrics"`
	Potential          float64       `json:"potential"`
	ReservationAllowed *bool         `json:"reservationAllowed,omitempty"`
	ReservedAt         *time.Time    `json:"reservedAt,omitempty"`
	ScannedAt          *time.Time    `json:"scannedAt,omitempty"`
	Screenshot         string        `json:"screenshot"`
	SocialMedia        []SocialMedia `json:"socialMedia"`
	Technologies       []Technology  `json:"technologies"`
	UnlockedAt         time.Time     `json:"unlockedAt"`
}

func (o *UnlockedLead) Validate() error {
	if o.BusinessFields == nil {
		return errors.New("property businessFields is required, but not set")
	}
	if err := o.Contact.Validate(); err != nil {
		return fmt.Errorf("invalid property contact: %w", err)
	}
	if err := o.Hoster.Validate(); err != nil {
		return fmt.Errorf("invalid property hoster: %w", err)
	}
	if o.Languages == nil {
		return errors.New("property languages is required, but not set")
	}
	if err := func() error {
		if o.MainTechnology == nil {
			return nil
		}
		return o.MainTechnology.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property mainTechnology: %w", err)
	}
	if err := o.Metrics.Validate(); err != nil {
		return fmt.Errorf("invalid property metrics: %w", err)
	}
	if o.SocialMedia == nil {
		return errors.New("property socialMedia is required, but not set")
	}
	if err := func() error {
		for i := range o.SocialMedia {
			if err := o.SocialMedia[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property socialMedia: %w", err)
	}
	if o.Technologies == nil {
		return errors.New("property technologies is required, but not set")
	}
	if err := func() error {
		for i := range o.Technologies {
			if err := o.Technologies[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property technologies: %w", err)
	}
	return nil
}
