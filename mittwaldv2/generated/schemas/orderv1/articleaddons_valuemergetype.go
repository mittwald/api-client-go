package orderv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "add"
//    - "set"

type ArticleAddonsValueMergeType string

const ArticleAddonsValueMergeTypeAdd ArticleAddonsValueMergeType = "add"
const ArticleAddonsValueMergeTypeSet ArticleAddonsValueMergeType = "set"

func (e ArticleAddonsValueMergeType) Validate() error {
	if e == ArticleAddonsValueMergeTypeAdd || e == ArticleAddonsValueMergeTypeSet {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
