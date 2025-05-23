package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "extension"

type RelatedAggregateReferenceAlternative6Domain string

const RelatedAggregateReferenceAlternative6DomainExtension RelatedAggregateReferenceAlternative6Domain = "extension"

func (e RelatedAggregateReferenceAlternative6Domain) Validate() error {
	if e == RelatedAggregateReferenceAlternative6DomainExtension {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
