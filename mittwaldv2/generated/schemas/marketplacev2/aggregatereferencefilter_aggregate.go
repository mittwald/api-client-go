package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "project"
//    - "customer"

type AggregateReferenceFilterAggregate string

const AggregateReferenceFilterAggregateProject AggregateReferenceFilterAggregate = "project"
const AggregateReferenceFilterAggregateCustomer AggregateReferenceFilterAggregate = "customer"

func (e AggregateReferenceFilterAggregate) Validate() error {
	if e == AggregateReferenceFilterAggregateProject || e == AggregateReferenceFilterAggregateCustomer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
