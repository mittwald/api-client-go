package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "aggregate":
//        type: "string"
//        enum:
//            - "project"
//            - "customer"
//    "domain":
//        type: "string"
//        enum:
//            - "project"
//            - "customer"
//    "id":
//        type: "string"
// required:
//    - "id"
//    - "domain"
//    - "aggregate"

type AggregateReference struct {
	Aggregate AggregateReferenceAggregate `json:"aggregate"`
	Domain    AggregateReferenceDomain    `json:"domain"`
	Id        string                      `json:"id"`
}

func (o *AggregateReference) Validate() error {
	if err := o.Aggregate.Validate(); err != nil {
		return fmt.Errorf("invalid property aggregate: %w", err)
	}
	if err := o.Domain.Validate(); err != nil {
		return fmt.Errorf("invalid property domain: %w", err)
	}
	return nil
}
