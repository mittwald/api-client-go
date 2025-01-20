package signupv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "api_read"
//    - "api_write"

type ApiTokenRolesItem string

const ApiTokenRolesItemAPIread ApiTokenRolesItem = "api_read"
const ApiTokenRolesItemAPIwrite ApiTokenRolesItem = "api_write"

func (e ApiTokenRolesItem) Validate() error {
	if e == ApiTokenRolesItemAPIread || e == ApiTokenRolesItemAPIwrite {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}