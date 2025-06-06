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
//    "businessFields":
//        type: "array"
//        items:
//            type: "string"
//    "company": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.BasicCompany"}
//    "description":
//        type: "string"
//    "hoster":
//        type: "object"
//        properties:
//            "server":
//                type: "array"
//                items:
//                    type: "string"
//        required:
//            - "server"
//    "languages":
//        type: "array"
//        items:
//            type: "string"
//    "leadId":
//        type: "string"
//    "mainTechnology": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Technology"}
//    "metrics": {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.BasicMetrics"}
//    "potential":
//        type: "number"
//        maximum: 1
//        minimum: 0
//        format: "float"
//    "scannedAt":
//        type: "string"
//        format: "date-time"
//    "screenshot":
//        type: "string"
//    "technologies":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.leadfyndr.Technology"}
// required:
//    - "leadId"
//    - "potential"
//    - "screenshot"
//    - "company"
//    - "metrics"
//    - "businessFields"
//    - "description"
//    - "technologies"
//    - "hoster"
//    - "languages"

type Lead struct {
	BusinessFields []string     `json:"businessFields"`
	Company        BasicCompany `json:"company"`
	Description    string       `json:"description"`
	Hoster         LeadHoster   `json:"hoster"`
	Languages      []string     `json:"languages"`
	LeadId         string       `json:"leadId"`
	MainTechnology *Technology  `json:"mainTechnology,omitempty"`
	Metrics        BasicMetrics `json:"metrics"`
	Potential      float64      `json:"potential"`
	ScannedAt      *time.Time   `json:"scannedAt,omitempty"`
	Screenshot     string       `json:"screenshot"`
	Technologies   []Technology `json:"technologies"`
}

func (o *Lead) Validate() error {
	if o.BusinessFields == nil {
		return errors.New("property businessFields is required, but not set")
	}
	if err := o.Company.Validate(); err != nil {
		return fmt.Errorf("invalid property company: %w", err)
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
