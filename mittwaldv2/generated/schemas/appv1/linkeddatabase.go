package appv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "databaseId":
//        type: "string"
//        format: "uuid"
//    "databaseUserIds":
//        type: "object"
//        additionalProperties:
//            type: "string"
//    "kind":
//        type: "string"
//        enum:
//            - "mysql"
//            - "redis"
//    "purpose":
//        type: "string"
//        enum:
//            - "primary"
//            - "cache"
//            - "custom"
//required:
//    - "databaseId"
//    - "purpose"
//    - "kind"
//description: "LinkedDatabase is a reference to a concrete Database and DatabaseUsers."

//LinkedDatabase is a reference to a concrete Database and DatabaseUsers.
type LinkedDatabase struct {
	DatabaseId      uuid.UUID             `json:"databaseId"`
	DatabaseUserIds map[string]string     `json:"databaseUserIds,omitempty"`
	Kind            LinkedDatabaseKind    `json:"kind"`
	Purpose         LinkedDatabasePurpose `json:"purpose"`
}

func (o *LinkedDatabase) Validate() error {
	if err := o.Kind.Validate(); err != nil {
		return fmt.Errorf("invalid property kind: %w", err)
	}
	if err := o.Purpose.Validate(); err != nil {
		return fmt.Errorf("invalid property purpose: %w", err)
	}
	return nil
}
