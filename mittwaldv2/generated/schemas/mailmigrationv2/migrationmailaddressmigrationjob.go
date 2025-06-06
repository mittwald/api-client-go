package mailmigrationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "migrate": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.MigrationMailAddressMigrationJobMigrate"}
// required:
//    - "migrate"

type MigrationMailAddressMigrationJob struct {
	Migrate MigrationMailAddressMigrationJobMigrate `json:"migrate"`
}

func (o *MigrationMailAddressMigrationJob) Validate() error {
	if err := o.Migrate.Validate(); err != nil {
		return fmt.Errorf("invalid property migrate: %w", err)
	}
	return nil
}
