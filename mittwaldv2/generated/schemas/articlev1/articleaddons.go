package articlev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "hidden":
//        type: "boolean"
//    "key":
//        type: "string"
//        minLength: 1
//    "type":
//        type: "string"
//    "value":
//        type: "string"
//    "valueMergeType":
//        type: "string"
//        enum:
//            - "add"
//            - "set"
//required:
//    - "key"
//    - "value"

//
type ArticleAddons struct {
	Hidden         *bool                        `json:"hidden,omitempty"`
	Key            string                       `json:"key"`
	Type           *string                      `json:"type,omitempty"`
	Value          string                       `json:"value"`
	ValueMergeType *ArticleAddonsValueMergeType `json:"valueMergeType,omitempty"`
}

func (o *ArticleAddons) Validate() error {
	if err := func() error {
		if o.ValueMergeType == nil {
			return nil
		}
		return o.ValueMergeType.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property valueMergeType: %w", err)
	}
	return nil
}