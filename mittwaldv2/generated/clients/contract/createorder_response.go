package contract

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "orderId":
//        type: "string"
//required:
//    - "orderId"

//
type CreateOrderResponse struct {
	OrderId string `json:"orderId"`
}

func (o *CreateOrderResponse) Validate() error {
	return nil
}
