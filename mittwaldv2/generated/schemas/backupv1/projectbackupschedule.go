package backupv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "description":
//        type: "string"
//        description: "Description of this ProjectBackupSchedule."
//        example: "I'm a ProjectBackupSchedule"
//    "id":
//        type: "string"
//        format: "uuid"
//        description: "ID of this ProjectBackupSchedule."
//    "isSystemBackup":
//        type: "boolean"
//        description: "True if this ProjectBackupSchedule was created and is managed by mittwald."
//    "projectId":
//        type: "string"
//        format: "uuid"
//        description: "ID of the Project this ProjectBackupSchedule belongs to."
//    "schedule":
//        type: "string"
//        description: "Execution schedule in crontab notation."
//        example: "5 4 * * *"
//    "ttl":
//        type: "string"
//        description: "TTL of the ProjectBackupSchedule as time string."
//        example: "7d"
//    "updatedAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "id"
//    - "projectId"
//    - "schedule"
//    - "isSystemBackup"

//
type ProjectBackupSchedule struct {
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Id             uuid.UUID  `json:"id"`
	IsSystemBackup bool       `json:"isSystemBackup"`
	ProjectId      uuid.UUID  `json:"projectId"`
	Schedule       string     `json:"schedule"`
	Ttl            *string    `json:"ttl,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

func (o *ProjectBackupSchedule) Validate() error {
	return nil
}