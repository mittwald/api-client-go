package feev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "feeStrategy": {"$ref": "#/components/schemas/de.mittwald.v1.fee.FeeStrategy"}
//    "id":
//        type: "string"
//        description: "The id of the given Resource"
//        example: "1241a1de-b633-4e6e-9aa5-7e878ccd6864"
//required:
//    - "id"
//description: "A Fee of a Resource"

// A Fee of a Resource
type ResourceFee struct {
	FeeStrategy *FeeStrategy `json:"feeStrategy,omitempty"`
	Id          string       `json:"id"`
}

func (o *ResourceFee) Validate() error {
	if err := func() error {
		if o.FeeStrategy == nil {
			return nil
		}
		return o.FeeStrategy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property feeStrategy: %w", err)
	}
	return nil
}
