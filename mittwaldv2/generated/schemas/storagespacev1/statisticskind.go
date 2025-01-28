package storagespacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "server"
//    - "project"

type StatisticsKind string

const StatisticsKindServer StatisticsKind = "server"
const StatisticsKindProject StatisticsKind = "project"

func (e StatisticsKind) Validate() error {
	if e == StatisticsKindServer || e == StatisticsKindProject {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
