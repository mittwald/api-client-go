package orderv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "nonRecurring":
//        type: "number"
//        example: 1000
//    "recurring":
//        type: "number"
//        example: 9000
//    "summary":
//        type: "number"
//        description: "The total price of the order"
//        example: 10000
//required:
//    - "summary"
//    - "recurring"
//    - "nonRecurring"

//
type OrderSummary struct {
	NonRecurring float64 `json:"nonRecurring"`
	Recurring    float64 `json:"recurring"`
	Summary      float64 `json:"summary"`
}

func (o *OrderSummary) Validate() error {
	return nil
}
