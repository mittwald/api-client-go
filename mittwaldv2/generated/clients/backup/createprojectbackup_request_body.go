package backup

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "description":
//        type: "string"
//        description: "Description of the Backup."
//        example: "I'm a ProjectBackup"
//    "expirationTime":
//        type: "string"
//        format: "date-time"
//        description: "Time when to expire the Backup."
//required:
//    - "expirationTime"

//
type CreateProjectBackupRequestBody struct {
	Description    *string   `json:"description,omitempty"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func (o *CreateProjectBackupRequestBody) Validate() error {
	return nil
}
