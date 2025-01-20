package conversationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "aggregate":
//        type: "string"
//        enum:
//            - "placementgroup"
//    "domain":
//        type: "string"
//        enum:
//            - "project"
//    "id":
//        type: "string"
//required:
//    - "id"
//    - "aggregate"
//    - "domain"

//
type ShareableAggregateReferenceAlternative4 struct {
	Aggregate ShareableAggregateReferenceAlternative4Aggregate `json:"aggregate"`
	Domain    ShareableAggregateReferenceAlternative4Domain    `json:"domain"`
	Id        string                                           `json:"id"`
}

func (o *ShareableAggregateReferenceAlternative4) Validate() error {
	if err := o.Aggregate.Validate(); err != nil {
		return fmt.Errorf("invalid property aggregate: %w", err)
	}
	if err := o.Domain.Validate(); err != nil {
		return fmt.Errorf("invalid property domain: %w", err)
	}
	return nil
}
