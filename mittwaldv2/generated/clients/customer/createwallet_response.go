package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "customerId":
//        type: "string"
//        example: "edefeee4-e8e9-4e2d-ab95-9a2eb6104cfb"

type CreateWalletResponse struct {
	CustomerId *string `json:"customerId,omitempty"`
}

func (o *CreateWalletResponse) Validate() error {
	return nil
}
