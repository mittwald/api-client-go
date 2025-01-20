package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
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
//    "isCancelled":
//        type: "boolean"

//
type CancelContractItemTerminationResponse struct {
	ContractId     *uuid.UUID `json:"contractId,omitempty"`
	ContractItemId *uuid.UUID `json:"contractItemId,omitempty"`
	IsCancelled    *bool      `json:"isCancelled,omitempty"`
}

func (o *CancelContractItemTerminationResponse) Validate() error {
	return nil
}