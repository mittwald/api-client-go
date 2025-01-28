package app

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "databaseId":
//        type: "string"
//        format: "uuid"
//    "databaseUserIds":
//        type: "object"
//        additionalProperties:
//            type: "string"
//    "purpose":
//        type: "string"
//        enum:
//            - "primary"
//            - "cache"
//            - "custom"
// required:
//    - "databaseId"
//    - "purpose"
// description: DeprecatedLinkDatabaseRequestBody models the JSON body of a 'deprecated-app-link-database' request

// DeprecatedLinkDatabaseRequestBody models the JSON body of a 'deprecated-app-link-database' request
type DeprecatedLinkDatabaseRequestBody struct {
	DatabaseId      string                                   `json:"databaseId"`
	DatabaseUserIds map[string]string                        `json:"databaseUserIds,omitempty"`
	Purpose         DeprecatedLinkDatabaseRequestBodyPurpose `json:"purpose"`
}

func (o *DeprecatedLinkDatabaseRequestBody) Validate() error {
	if err := o.Purpose.Validate(); err != nil {
		return fmt.Errorf("invalid property purpose: %w", err)
	}
	return nil
}
