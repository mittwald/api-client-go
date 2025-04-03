package backupclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//        description: "Description of the ProjectBackupSchedule. Note that the description of isSystemBackup true items cannot be changed."
//        example: "I'm a ProjectBackupSchedule"
//    "schedule":
//        type: "string"
//        description: "Execution schedule in crontab notation. Note that the schedule of isSystemBackup true items must be daily once."
//        example: "5 4 * * *"
//    "ttl":
//        type: "string"
//        description: "TTL of the ProjectBackupSchedule as time string."
//        example: "7d"
// description: UpdateProjectBackupScheduleRequestBody models the JSON body of a 'backup-update-project-backup-schedule' request

// UpdateProjectBackupScheduleRequestBody models the JSON body of a 'backup-update-project-backup-schedule' request
type UpdateProjectBackupScheduleRequestBody struct {
	Description *string `json:"description,omitempty"`
	Schedule    *string `json:"schedule,omitempty"`
	Ttl         *string `json:"ttl,omitempty"`
}

func (o *UpdateProjectBackupScheduleRequestBody) Validate() error {
	return nil
}
