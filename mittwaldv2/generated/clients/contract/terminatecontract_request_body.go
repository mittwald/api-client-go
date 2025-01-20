package contract

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "reason":
//        type: "string"
//        description: "A reason for the termination can be given as plain text."
//        example: "Projekt wird nicht mehr benötigt"
//    "terminationTargetDate":
//        type: "string"
//        format: "date-time"
//        description: "The termination date has to be a valid date according to activation and contract period of the base ContractItem. If none given, the next possible termination date will be used."

//
type TerminateContractRequestBody struct {
	Reason                *string    `json:"reason,omitempty"`
	TerminationTargetDate *time.Time `json:"terminationTargetDate,omitempty"`
}

func (o *TerminateContractRequestBody) Validate() error {
	return nil
}