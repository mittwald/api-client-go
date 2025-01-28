package domain

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "jsonSchemaAdminC":
//        type: "object"
//    "jsonSchemaOwnerC":
//        type: "object"
// required:
//    - "jsonSchemaOwnerC"

type DeprecatedGetHandleFieldsResponse struct {
	JsonSchemaAdminC *DeprecatedGetHandleFieldsResponseJsonSchemaAdminC `json:"jsonSchemaAdminC,omitempty"`
	JsonSchemaOwnerC DeprecatedGetHandleFieldsResponseJsonSchemaOwnerC  `json:"jsonSchemaOwnerC"`
}

func (o *DeprecatedGetHandleFieldsResponse) Validate() error {
	if err := func() error {
		if o.JsonSchemaAdminC == nil {
			return nil
		}
		return o.JsonSchemaAdminC.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property jsonSchemaAdminC: %w", err)
	}
	if err := o.JsonSchemaOwnerC.Validate(); err != nil {
		return fmt.Errorf("invalid property jsonSchemaOwnerC: %w", err)
	}
	return nil
}
