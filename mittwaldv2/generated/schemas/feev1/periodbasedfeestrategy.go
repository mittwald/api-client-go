package feev1

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "periods":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "feeValidFrom":
//                    type: "string"
//                    format: "date-time"
//                "feeValidUntil":
//                    type: "string"
//                    format: "date-time"
//                "monthlyPrice":
//                    type: "number"
//                    description: "The monthly price in Euro Cents."
//            required:
//                - "monthlyPrice"
//required:
//    - "periods"
//description: "A strategy for fees that occur periodically"

//A strategy for fees that occur periodically
type PeriodBasedFeeStrategy struct {
	Periods []PeriodBasedFeeStrategyPeriodsItem `json:"periods"`
}

func (o *PeriodBasedFeeStrategy) Validate() error {
	if o.Periods == nil {
		return errors.New("property periods is required, but not set")
	}
	if err := func() error {
		for i := range o.Periods {
			if err := o.Periods[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property periods: %w", err)
	}
	return nil
}
