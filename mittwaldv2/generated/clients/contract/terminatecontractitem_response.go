package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//properties:
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
	ContractId            *uuid.UUID `json:"contractId,omitempty"`
	ContractItemId        *uuid.UUID `json:"contractItemId,omitempty"`
	Reason                *string    `json:"reason,omitempty"`
	TerminationTargetDate *time.Time `json:"terminationTargetDate,omitempty"`
}

func (o *TerminateContractItemResponse) Validate() error {
	return nil
}
