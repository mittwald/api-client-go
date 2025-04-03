package projectv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "creating"
//    - "ready"
//    - "unready"
// description: "deprecated by property status"
// deprecated: true

// deprecated by property status
type DeprecatedServerReadinessStatus string

const DeprecatedServerReadinessStatusCreating DeprecatedServerReadinessStatus = "creating"
const DeprecatedServerReadinessStatusReady DeprecatedServerReadinessStatus = "ready"
const DeprecatedServerReadinessStatusUnready DeprecatedServerReadinessStatus = "unready"

func (e DeprecatedServerReadinessStatus) Validate() error {
	if e == DeprecatedServerReadinessStatusCreating || e == DeprecatedServerReadinessStatusReady || e == DeprecatedServerReadinessStatusUnready {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
