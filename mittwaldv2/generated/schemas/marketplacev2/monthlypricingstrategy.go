package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "netPrice":
//        type: "integer"
//        format: "int32"
//        description: "The monthly price in Euro Cents before tax."
//        example: 500
// required:
//    - "netPrice"
// description: "A strategy for pricing that occurs monthly."

// A strategy for pricing that occurs monthly.
type MonthlyPricingStrategy struct {
	NetPrice int64 `json:"netPrice"`
}

func (o *MonthlyPricingStrategy) Validate() error {
	return nil
}
