package mailmigrationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "requirements": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.MigrateMailboxCommandRequirements"}
//required:
//    - "requirements"

//
type MigrationMailboxJobMigrate struct {
	Requirements MigrateMailboxCommandRequirements `json:"requirements"`
}

func (o *MigrationMailboxJobMigrate) Validate() error {
	if err := o.Requirements.Validate(); err != nil {
		return fmt.Errorf("invalid property requirements: %w", err)
	}
	return nil
}
