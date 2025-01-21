package relocation

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "positions":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "name":
//                    type: "string"
//                    minLength: 1
//                "price":
//                    type: "number"
//                    minimum: 0
//            required:
//                - "name"
//                - "price"
//    "total":
//        type: "number"
//        minimum: 0
//required:
//    - "positions"
//    - "total"

type CreateRelocationRequestBodyPrices struct {
	Positions []CreateRelocationRequestBodyPricesPositionsItem `json:"positions"`
	Total     float64                                          `json:"total"`
}

func (o *CreateRelocationRequestBodyPrices) Validate() error {
	if o.Positions == nil {
		return errors.New("property positions is required, but not set")
	}
	if err := func() error {
		for i := range o.Positions {
			if err := o.Positions[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property positions: %w", err)
	}
	return nil
}
