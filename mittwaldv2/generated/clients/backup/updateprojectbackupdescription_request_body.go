package backup

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//        description: "Description of the ProjectBackup."
//        example: "I'm a ProjectBackup"
// description: UpdateProjectBackupDescriptionRequestBody models the JSON body of a 'backup-update-project-backup-description' request

// UpdateProjectBackupDescriptionRequestBody models the JSON body of a 'backup-update-project-backup-description' request
type UpdateProjectBackupDescriptionRequestBody struct {
	Description *string `json:"description,omitempty"`
}

func (o *UpdateProjectBackupDescriptionRequestBody) Validate() error {
	return nil
}
