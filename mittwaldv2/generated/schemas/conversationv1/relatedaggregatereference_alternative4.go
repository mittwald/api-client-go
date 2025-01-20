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
//            - "appinstallation"
//    "domain":
//        type: "string"
//        enum:
//            - "app"
//    "id":
//        type: "string"
//required:
//    - "id"
//    - "aggregate"
//    - "domain"

//
type RelatedAggregateReferenceAlternative4 struct {
	Aggregate RelatedAggregateReferenceAlternative4Aggregate `json:"aggregate"`
	Domain    RelatedAggregateReferenceAlternative4Domain    `json:"domain"`
	Id        string                                         `json:"id"`
}

func (o *RelatedAggregateReferenceAlternative4) Validate() error {
	if err := o.Aggregate.Validate(); err != nil {
		return fmt.Errorf("invalid property aggregate: %w", err)
	}
	if err := o.Domain.Validate(); err != nil {
		return fmt.Errorf("invalid property domain: %w", err)
	}
	return nil
}