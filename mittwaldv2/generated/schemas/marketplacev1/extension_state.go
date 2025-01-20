package marketplacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "enabled"
//    - "blocked"
//    - "disabled"
//description: "deprecated"
//deprecated: true

//deprecated
type ExtensionState string

const ExtensionStateEnabled ExtensionState = "enabled"
const ExtensionStateBlocked ExtensionState = "blocked"
const ExtensionStateDisabled ExtensionState = "disabled"

func (e ExtensionState) Validate() error {
	if e == ExtensionStateEnabled || e == ExtensionStateBlocked || e == ExtensionStateDisabled {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
