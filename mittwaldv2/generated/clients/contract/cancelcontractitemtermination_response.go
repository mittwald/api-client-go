package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// properties:
//    "contractId":
//        type: "string"
//        format: "uuid"
//    "contractItemId":
//        type: "string"
//        format: "uuid"
//    "isCancelled":
//        type: "boolean"

type CancelContractItemTerminationResponse struct {
	ContractId     *string `json:"contractId,omitempty"`
	ContractItemId *string `json:"contractItemId,omitempty"`
	IsCancelled    *bool   `json:"isCancelled,omitempty"`
}

func (o *CancelContractItemTerminationResponse) Validate() error {
	return nil
}
