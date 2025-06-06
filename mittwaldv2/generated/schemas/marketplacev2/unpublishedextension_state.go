package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "enabled"
//    - "blocked"
//    - "disabled"
// description: "deprecated"
// deprecated: true

// deprecated
type UnpublishedExtensionState string

const UnpublishedExtensionStateEnabled UnpublishedExtensionState = "enabled"
const UnpublishedExtensionStateBlocked UnpublishedExtensionState = "blocked"
const UnpublishedExtensionStateDisabled UnpublishedExtensionState = "disabled"

func (e UnpublishedExtensionState) Validate() error {
	if e == UnpublishedExtensionStateEnabled || e == UnpublishedExtensionStateBlocked || e == UnpublishedExtensionStateDisabled {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
