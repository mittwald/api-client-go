package databasev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessIpMask":
//        type: "string"
//    "accessLevel":
//        type: "string"
//        enum:
//            - "full"
//            - "readonly"
//    "databaseId":
//        type: "string"
//        format: "uuid"
//    "description":
//        type: "string"
//    "externalAccess":
//        type: "boolean"
//    "password":
//        type: "string"
//required:
//    - "databaseId"
//    - "password"
//    - "description"
//    - "accessLevel"

//
type CreateMySqlUser struct {
	AccessIpMask   *string                    `json:"accessIpMask,omitempty"`
	AccessLevel    CreateMySqlUserAccessLevel `json:"accessLevel"`
	DatabaseId     uuid.UUID                  `json:"databaseId"`
	Description    string                     `json:"description"`
	ExternalAccess *bool                      `json:"externalAccess,omitempty"`
	Password       string                     `json:"password"`
}

func (o *CreateMySqlUser) Validate() error {
	if err := o.AccessLevel.Validate(); err != nil {
		return fmt.Errorf("invalid property accessLevel: %w", err)
	}
	return nil
}
