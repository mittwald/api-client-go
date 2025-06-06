package orderv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "other"
//    - "mr"
//    - "ms"
// description: "the users title"
// deprecated: true

// the users title
type ProfileTitle string

const ProfileTitleOther ProfileTitle = "other"
const ProfileTitleMr ProfileTitle = "mr"
const ProfileTitleMs ProfileTitle = "ms"

func (e ProfileTitle) Validate() error {
	if e == ProfileTitleOther || e == ProfileTitleMr || e == ProfileTitleMs {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
