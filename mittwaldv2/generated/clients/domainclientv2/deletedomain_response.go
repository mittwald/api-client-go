package domainclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "isAsync":
//        type: "boolean"
//    "transactionId":
//        type: "string"

type DeleteDomainResponse struct {
	IsAsync       *bool   `json:"isAsync,omitempty"`
	TransactionId *string `json:"transactionId,omitempty"`
}

func (o *DeleteDomainResponse) Validate() error {
	return nil
}
