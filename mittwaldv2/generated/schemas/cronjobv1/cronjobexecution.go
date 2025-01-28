package cronjobv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "abortedBy":
//        type: "object"
//        properties:
//            "id":
//                type: "string"
//                format: "uuid"
//    "cronjobId":
//        type: "string"
//    "durationInMilliseconds":
//        type: "integer"
//        format: "int64"
//        example: 12374
//    "end":
//        type: "string"
//        format: "date-time"
//    "executionEnd":
//        type: "string"
//        format: "date-time"
//        deprecated: true
//    "executionStart":
//        type: "string"
//        format: "date-time"
//        deprecated: true
//    "id":
//        type: "string"
//        example: "cron-bd26li-28027320"
//    "logPath":
//        type: "string"
//        example: "/var/log/cronjobs/cron-bd26li-28027320.log"
//    "start":
//        type: "string"
//        format: "date-time"
//    "status":
//        type: "string"
//        enum:
//            - "Complete"
//            - "Failed"
//            - "AbortedBySystem"
//            - "Pending"
//            - "Running"
//            - "AbortedByUser"
//            - "TimedOut"
//    "successful":
//        type: "boolean"
//    "triggeredBy":
//        type: "object"
//        properties:
//            "id":
//                type: "string"
//                format: "uuid"
// required:
//    - "id"
//    - "status"
//    - "successful"
//    - "cronjobId"

type CronjobExecution struct {
	AbortedBy              *CronjobExecutionAbortedBy   `json:"abortedBy,omitempty"`
	CronjobId              string                       `json:"cronjobId"`
	DurationInMilliseconds *int64                       `json:"durationInMilliseconds,omitempty"`
	End                    *time.Time                   `json:"end,omitempty"`
	ExecutionEnd           *time.Time                   `json:"executionEnd,omitempty"`
	ExecutionStart         *time.Time                   `json:"executionStart,omitempty"`
	Id                     string                       `json:"id"`
	LogPath                *string                      `json:"logPath,omitempty"`
	Start                  *time.Time                   `json:"start,omitempty"`
	Status                 CronjobExecutionStatus       `json:"status"`
	Successful             bool                         `json:"successful"`
	TriggeredBy            *CronjobExecutionTriggeredBy `json:"triggeredBy,omitempty"`
}

func (o *CronjobExecution) Validate() error {
	if err := func() error {
		if o.AbortedBy == nil {
			return nil
		}
		return o.AbortedBy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property abortedBy: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	if err := func() error {
		if o.TriggeredBy == nil {
			return nil
		}
		return o.TriggeredBy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property triggeredBy: %w", err)
	}
	return nil
}
