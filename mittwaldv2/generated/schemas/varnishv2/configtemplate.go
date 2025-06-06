package varnishv2

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
//    "files":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.varnish.ConfigFileRef"}
//        description: "A set of config file references."
//    "isGlobal":
//        type: "boolean"
//    "note":
//        type: "string"
//    "projectId":
//        type: "string"
//    "softwareConfigTemplateId":
//        type: "string"
//    "softwareTemplateId":
//        type: "string"
//    "updatedAt":
//        type: "string"
//        format: "date-time"
// required:
//    - "softwareConfigTemplateId"
//    - "softwareTemplateId"
//    - "files"

type ConfigTemplate struct {
	Files                    []ConfigFileRef `json:"files"`
	IsGlobal                 *bool           `json:"isGlobal,omitempty"`
	Note                     *string         `json:"note,omitempty"`
	ProjectId                *string         `json:"projectId,omitempty"`
	SoftwareConfigTemplateId string          `json:"softwareConfigTemplateId"`
	SoftwareTemplateId       string          `json:"softwareTemplateId"`
	UpdatedAt                *time.Time      `json:"updatedAt,omitempty"`
}

func (o *ConfigTemplate) Validate() error {
	if o.Files == nil {
		return errors.New("property files is required, but not set")
	}
	if err := func() error {
		for i := range o.Files {
			if err := o.Files[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property files: %w", err)
	}
	return nil
}
