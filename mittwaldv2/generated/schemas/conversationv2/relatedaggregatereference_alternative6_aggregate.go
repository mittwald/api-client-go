package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "extensionInstance"

type RelatedAggregateReferenceAlternative6Aggregate string

const RelatedAggregateReferenceAlternative6AggregateExtensionInstance RelatedAggregateReferenceAlternative6Aggregate = "extensionInstance"

func (e RelatedAggregateReferenceAlternative6Aggregate) Validate() error {
	if e == RelatedAggregateReferenceAlternative6AggregateExtensionInstance {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
