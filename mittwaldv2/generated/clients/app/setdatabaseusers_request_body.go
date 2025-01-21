package app

import "errors"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "databaseUserIds":
//        type: "object"
//        additionalProperties:
//            type: "string"
//required:
//    - "databaseUserIds"

type SetDatabaseUsersRequestBody struct {
	DatabaseUserIds map[string]string `json:"databaseUserIds"`
}

func (o *SetDatabaseUsersRequestBody) Validate() error {
	if o.DatabaseUserIds == nil {
		return errors.New("property databaseUserIds is required, but not set")
	}
	return nil
}
