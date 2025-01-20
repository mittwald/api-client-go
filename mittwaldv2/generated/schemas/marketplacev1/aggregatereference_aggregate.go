package marketplacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "project"
//    - "customer"

type AggregateReferenceAggregate string

const AggregateReferenceAggregateProject AggregateReferenceAggregate = "project"
const AggregateReferenceAggregateCustomer AggregateReferenceAggregate = "customer"

func (e AggregateReferenceAggregate) Validate() error {
	if e == AggregateReferenceAggregateProject || e == AggregateReferenceAggregateCustomer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}