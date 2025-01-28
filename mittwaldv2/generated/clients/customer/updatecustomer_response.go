package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "customerId":
//        type: "string"
//    "customerNumber":
//        type: "string"
//    "name":
//        type: "string"
// required:
//    - "customerId"
//    - "name"
//    - "customerNumber"

type UpdateCustomerResponse struct {
	CustomerId     string `json:"customerId"`
	CustomerNumber string `json:"customerNumber"`
	Name           string `json:"name"`
}

func (o *UpdateCustomerResponse) Validate() error {
	return nil
}
