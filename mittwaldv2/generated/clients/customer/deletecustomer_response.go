package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "customerId":
//        type: "string"

type DeleteCustomerResponse struct {
	CustomerId *string `json:"customerId,omitempty"`
}

func (o *DeleteCustomerResponse) Validate() error {
	return nil
}
