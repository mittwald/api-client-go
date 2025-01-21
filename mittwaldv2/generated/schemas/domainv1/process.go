package domainv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "error":
//        type: "string"
//    "lastUpdate":
//        type: "string"
//        format: "date-time"
//    "processType": {"$ref": "#/components/schemas/de.mittwald.v1.domain.ProcessType"}
//    "state": {"$ref": "#/components/schemas/de.mittwald.v1.domain.ProcessState"}
//    "status":
//        type: "string"
//    "statusCode":
//        type: "string"
//    "transactionId":
//        type: "string"
//required:
//    - "transactionId"
//    - "processType"
//    - "state"
//    - "lastUpdate"

type Process struct {
	Error         *string      `json:"error,omitempty"`
	LastUpdate    time.Time    `json:"lastUpdate"`
	ProcessType   ProcessType  `json:"processType"`
	State         ProcessState `json:"state"`
	Status        *string      `json:"status,omitempty"`
	StatusCode    *string      `json:"statusCode,omitempty"`
	TransactionId string       `json:"transactionId"`
}

func (o *Process) Validate() error {
	if err := o.ProcessType.Validate(); err != nil {
		return fmt.Errorf("invalid property processType: %w", err)
	}
	if err := o.State.Validate(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	return nil
}
