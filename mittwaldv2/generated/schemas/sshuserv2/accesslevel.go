package sshuserv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "full"
//    - "read"
//    - "unspecified"
// description: "The level of access for an SFTPUser."

// The level of access for an SFTPUser.
type AccessLevel string

const AccessLevelFull AccessLevel = "full"
const AccessLevelRead AccessLevel = "read"
const AccessLevelUnspecified AccessLevel = "unspecified"

func (e AccessLevel) Validate() error {
	if e == AccessLevelFull || e == AccessLevelRead || e == AccessLevelUnspecified {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
