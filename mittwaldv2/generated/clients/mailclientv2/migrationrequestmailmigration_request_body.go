package mailclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "sourceLegacyProjectId":
//        type: "string"
//    "targetProjectId":
//        type: "string"
//        format: "uuid"
// required:
//    - "sourceLegacyProjectId"
//    - "targetProjectId"
// description: MigrationRequestMailMigrationRequestBody models the JSON body of a 'mail-migration-request-mail-migration' request

// MigrationRequestMailMigrationRequestBody models the JSON body of a 'mail-migration-request-mail-migration' request
type MigrationRequestMailMigrationRequestBody struct {
	SourceLegacyProjectId string `json:"sourceLegacyProjectId"`
	TargetProjectId       string `json:"targetProjectId"`
}

func (o *MigrationRequestMailMigrationRequestBody) Validate() error {
	return nil
}
