package conversationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "placementgroup"

type ShareableAggregateReferenceAlternative4Aggregate string

const ShareableAggregateReferenceAlternative4AggregatePlacementgroup ShareableAggregateReferenceAlternative4Aggregate = "placementgroup"

func (e ShareableAggregateReferenceAlternative4Aggregate) Validate() error {
	if e == ShareableAggregateReferenceAlternative4AggregatePlacementgroup {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}