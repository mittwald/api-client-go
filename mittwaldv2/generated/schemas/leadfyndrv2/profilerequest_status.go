package leadfyndrv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "AUTOTEST_INIT"
//    - "MANUAL_VERIFICATION"
//    - "REJECTED"
//    - "APPROVED"

type ProfileRequestStatus string

const ProfileRequestStatusAUTOTESTINIT ProfileRequestStatus = "AUTOTEST_INIT"
const ProfileRequestStatusMANUALVERIFICATION ProfileRequestStatus = "MANUAL_VERIFICATION"
const ProfileRequestStatusREJECTED ProfileRequestStatus = "REJECTED"
const ProfileRequestStatusAPPROVED ProfileRequestStatus = "APPROVED"

func (e ProfileRequestStatus) Validate() error {
	if e == ProfileRequestStatusAUTOTESTINIT || e == ProfileRequestStatusMANUALVERIFICATION || e == ProfileRequestStatusREJECTED || e == ProfileRequestStatusAPPROVED {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
