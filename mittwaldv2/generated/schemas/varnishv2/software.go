package varnishv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "config": {"$ref": "#/components/schemas/de.mittwald.v1.varnish.SoftwareConfig"}
//    "projectId":
//        type: "string"
//    "settings":
//        type: "object"
//        additionalProperties:
//            type: "string"
//    "softwareId":
//        type: "string"
//    "softwareTemplateId":
//        type: "string"
//    "softwareVersion":
//        type: "string"
//    "updatedAt":
//        type: "string"
//        format: "date-time"
// required:
//    - "softwareId"
//    - "updatedAt"
//    - "softwareTemplateId"
//    - "softwareVersion"
//    - "projectId"
//    - "config"

type Software struct {
	Config             SoftwareConfig    `json:"config"`
	ProjectId          string            `json:"projectId"`
	Settings           map[string]string `json:"settings,omitempty"`
	SoftwareId         string            `json:"softwareId"`
	SoftwareTemplateId string            `json:"softwareTemplateId"`
	SoftwareVersion    string            `json:"softwareVersion"`
	UpdatedAt          time.Time         `json:"updatedAt"`
}

func (o *Software) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return fmt.Errorf("invalid property config: %w", err)
	}
	return nil
}
