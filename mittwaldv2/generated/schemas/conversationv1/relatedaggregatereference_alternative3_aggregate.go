package conversationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "project"

type RelatedAggregateReferenceAlternative3Aggregate string

const RelatedAggregateReferenceAlternative3AggregateProject RelatedAggregateReferenceAlternative3Aggregate = "project"

func (e RelatedAggregateReferenceAlternative3Aggregate) Validate() error {
	if e == RelatedAggregateReferenceAlternative3AggregateProject {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
