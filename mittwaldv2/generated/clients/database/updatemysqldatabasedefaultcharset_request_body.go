package database

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/databasev1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// properties:
//    "characterSettings": {"$ref": "#/components/schemas/de.mittwald.v1.database.characterSettings"}
// required:
//    - "characterSettings"
// description: UpdateMysqlDatabaseDefaultCharsetRequestBody models the JSON body of a 'database-update-mysql-database-default-charset' request

// UpdateMysqlDatabaseDefaultCharsetRequestBody models the JSON body of a 'database-update-mysql-database-default-charset' request
type UpdateMysqlDatabaseDefaultCharsetRequestBody struct {
	CharacterSettings databasev1.CharacterSettings `json:"characterSettings"`
}

func (o *UpdateMysqlDatabaseDefaultCharsetRequestBody) Validate() error {
	if err := o.CharacterSettings.Validate(); err != nil {
		return fmt.Errorf("invalid property characterSettings: %w", err)
	}
	return nil
}
