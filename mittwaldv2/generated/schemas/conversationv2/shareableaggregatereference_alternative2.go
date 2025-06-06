package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "aggregate":
//        type: "string"
//        enum:
//            - "customer"
//    "domain":
//        type: "string"
//        enum:
//            - "customer"
//    "id":
//        type: "string"
// required:
//    - "id"
//    - "aggregate"
//    - "domain"

type ShareableAggregateReferenceAlternative2 struct {
	Aggregate ShareableAggregateReferenceAlternative2Aggregate `json:"aggregate"`
	Domain    ShareableAggregateReferenceAlternative2Domain    `json:"domain"`
	Id        string                                           `json:"id"`
}

func (o *ShareableAggregateReferenceAlternative2) Validate() error {
	if err := o.Aggregate.Validate(); err != nil {
		return fmt.Errorf("invalid property aggregate: %w", err)
	}
	if err := o.Domain.Validate(); err != nil {
		return fmt.Errorf("invalid property domain: %w", err)
	}
	return nil
}
