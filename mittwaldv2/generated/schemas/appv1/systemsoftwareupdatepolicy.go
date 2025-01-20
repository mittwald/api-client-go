package appv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "none"
//    - "inheritedFromApp"
//    - "patchLevel"
//    - "all"
//description: "SystemSoftwareUpdatePolicy describes which updates should be applied automatically by our systems."

//SystemSoftwareUpdatePolicy describes which updates should be applied automatically by our systems.
type SystemSoftwareUpdatePolicy string

const SystemSoftwareUpdatePolicyNone SystemSoftwareUpdatePolicy = "none"
const SystemSoftwareUpdatePolicyInheritedFromApp SystemSoftwareUpdatePolicy = "inheritedFromApp"
const SystemSoftwareUpdatePolicyPatchLevel SystemSoftwareUpdatePolicy = "patchLevel"
const SystemSoftwareUpdatePolicyAll SystemSoftwareUpdatePolicy = "all"

func (e SystemSoftwareUpdatePolicy) Validate() error {
	if e == SystemSoftwareUpdatePolicyNone || e == SystemSoftwareUpdatePolicyInheritedFromApp || e == SystemSoftwareUpdatePolicyPatchLevel || e == SystemSoftwareUpdatePolicyAll {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}