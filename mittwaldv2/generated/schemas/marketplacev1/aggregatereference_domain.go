package marketplacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "project"
//    - "customer"

type AggregateReferenceDomain string

const AggregateReferenceDomainProject AggregateReferenceDomain = "project"
const AggregateReferenceDomainCustomer AggregateReferenceDomain = "customer"

func (e AggregateReferenceDomain) Validate() error {
	if e == AggregateReferenceDomainProject || e == AggregateReferenceDomainCustomer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}