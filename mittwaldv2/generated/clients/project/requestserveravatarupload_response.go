package project

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/projectv1"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "refId":
//        type: "string"
//        format: "uuid"
//    "rules": {"$ref": "#/components/schemas/de.mittwald.v1.project.AvatarRules"}
//required:
//    - "refId"
//    - "rules"

//
type RequestServerAvatarUploadResponse struct {
	RefId uuid.UUID             `json:"refId"`
	Rules projectv1.AvatarRules `json:"rules"`
}

func (o *RequestServerAvatarUploadResponse) Validate() error {
	if err := o.Rules.Validate(); err != nil {
		return fmt.Errorf("invalid property rules: %w", err)
	}
	return nil
}
