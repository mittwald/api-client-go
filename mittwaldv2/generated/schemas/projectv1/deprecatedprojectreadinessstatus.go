package projectv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "creating"
//    - "ready"
//    - "unready"
//description: "deprecated by property status"
//deprecated: true

//deprecated by property status
type DeprecatedProjectReadinessStatus string

const DeprecatedProjectReadinessStatusCreating DeprecatedProjectReadinessStatus = "creating"
const DeprecatedProjectReadinessStatusReady DeprecatedProjectReadinessStatus = "ready"
const DeprecatedProjectReadinessStatusUnready DeprecatedProjectReadinessStatus = "unready"

func (e DeprecatedProjectReadinessStatus) Validate() error {
	if e == DeprecatedProjectReadinessStatusCreating || e == DeprecatedProjectReadinessStatusReady || e == DeprecatedProjectReadinessStatusUnready {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}