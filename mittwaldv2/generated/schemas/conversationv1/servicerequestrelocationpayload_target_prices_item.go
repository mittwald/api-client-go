package conversationv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "name":
//        type: "string"
//    "price":
//        type: "number"

//
type ServiceRequestRelocationPayloadTargetPricesItem struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
}

func (o *ServiceRequestRelocationPayloadTargetPricesItem) Validate() error {
	return nil
}
