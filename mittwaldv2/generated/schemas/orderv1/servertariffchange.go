package orderv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "contractId":
//        type: "string"
//        format: "uuid"
//    "diskspaceInGiB":
//        type: "number"
//        example: 100
//    "machineType":
//        type: "string"
//        example: "shared.xlarge"
//required:
//    - "contractId"
//    - "machineType"
//    - "diskspaceInGiB"

type ServerTariffChange struct {
	ContractId     uuid.UUID `json:"contractId"`
	DiskspaceInGiB float64   `json:"diskspaceInGiB"`
	MachineType    string    `json:"machineType"`
}

func (o *ServerTariffChange) Validate() error {
	return nil
}
