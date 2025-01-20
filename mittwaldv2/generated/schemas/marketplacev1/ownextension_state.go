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

//deprecated
type OwnExtensionState string

const OwnExtensionStateEnabled OwnExtensionState = "enabled"
const OwnExtensionStateBlocked OwnExtensionState = "blocked"
const OwnExtensionStateDisabled OwnExtensionState = "disabled"

func (e OwnExtensionState) Validate() error {
	if e == OwnExtensionStateEnabled || e == OwnExtensionStateBlocked || e == OwnExtensionStateDisabled {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}