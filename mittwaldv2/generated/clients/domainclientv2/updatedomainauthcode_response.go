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

type UpdateDomainAuthCodeResponse struct {
	IsAsync       *bool   `json:"isAsync,omitempty"`
	TransactionId *string `json:"transactionId,omitempty"`
}

func (o *UpdateDomainAuthCodeResponse) Validate() error {
	return nil
}
