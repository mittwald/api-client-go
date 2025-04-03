package contractclientv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "contractId":
//        type: "string"
//        format: "uuid"
//    "contractItemId":
//        type: "string"
//        format: "uuid"
//    "reason":
//        type: "string"
//        example: "Server wird nicht mehr benötigt"
//    "terminationTargetDate":
//        type: "string"
//        format: "date-time"

type TerminateContractItemResponse struct {
	ContractId            *string    `json:"contractId,omitempty"`
	ContractItemId        *string    `json:"contractItemId,omitempty"`
	Reason                *string    `json:"reason,omitempty"`
	TerminationTargetDate *time.Time `json:"terminationTargetDate,omitempty"`
}

func (o *TerminateContractItemResponse) Validate() error {
	return nil
}
