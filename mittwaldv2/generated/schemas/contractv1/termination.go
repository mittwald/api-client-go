package contractv1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "cancellationForbidden":
//        type: "boolean"
//        description: "Indicates whether the User is allowed to cancel the Termination."
//    "reason":
//        type: "string"
//        example: "Not needed anymore"
//    "scheduledAtDate":
//        type: "string"
//        format: "date-time"
//    "scheduledByUserId":
//        type: "string"
//    "targetDate":
//        type: "string"
//        format: "date-time"
// required:
//    - "scheduledAtDate"
//    - "targetDate"

type Termination struct {
	CancellationForbidden *bool     `json:"cancellationForbidden,omitempty"`
	Reason                *string   `json:"reason,omitempty"`
	ScheduledAtDate       time.Time `json:"scheduledAtDate"`
	ScheduledByUserId     *string   `json:"scheduledByUserId,omitempty"`
	TargetDate            time.Time `json:"targetDate"`
}

func (o *Termination) Validate() error {
	return nil
}
