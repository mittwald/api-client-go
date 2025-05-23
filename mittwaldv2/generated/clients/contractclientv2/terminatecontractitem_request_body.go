package contractclientv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "reason":
//        type: "string"
//        description: "A reason for the Termination can be given as plain text."
//        example: "Server wird nicht mehr benötigt"
//    "terminationTargetDate":
//        type: "string"
//        format: "date-time"
//        description: "The termination date has to be a valid date according to activation and contract period of the ContractItem. If none given, the next possible termination date will be used."
// description: TerminateContractItemRequestBody models the JSON body of a 'contract-terminate-contract-item' request

// TerminateContractItemRequestBody models the JSON body of a 'contract-terminate-contract-item' request
type TerminateContractItemRequestBody struct {
	Reason                *string    `json:"reason,omitempty"`
	TerminationTargetDate *time.Time `json:"terminationTargetDate,omitempty"`
}

func (o *TerminateContractItemRequestBody) Validate() error {
	return nil
}
