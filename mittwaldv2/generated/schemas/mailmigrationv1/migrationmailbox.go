package mailmigrationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//    "finished":
//        type: "boolean"
//    "id":
//        type: "string"
//        format: "uuid"
//    "migrationJobs": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.MigrationMailboxJob"}
//    "name":
//        type: "string"
// required:
//    - "id"
//    - "name"
//    - "migrationJobs"
//    - "finished"

type MigrationMailbox struct {
	Description   *string             `json:"description,omitempty"`
	Finished      bool                `json:"finished"`
	Id            string              `json:"id"`
	MigrationJobs MigrationMailboxJob `json:"migrationJobs"`
	Name          string              `json:"name"`
}

func (o *MigrationMailbox) Validate() error {
	if err := o.MigrationJobs.Validate(); err != nil {
		return fmt.Errorf("invalid property migrationJobs: %w", err)
	}
	return nil
}
