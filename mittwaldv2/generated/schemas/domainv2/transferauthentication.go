package domainv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "unspecified"
//    - "code"
//    - "email"
//    - "push"

type TransferAuthentication string

const TransferAuthenticationUnspecified TransferAuthentication = "unspecified"
const TransferAuthenticationCode TransferAuthentication = "code"
const TransferAuthenticationEmail TransferAuthentication = "email"
const TransferAuthenticationPush TransferAuthentication = "push"

func (e TransferAuthentication) Validate() error {
	if e == TransferAuthenticationUnspecified || e == TransferAuthenticationCode || e == TransferAuthenticationEmail || e == TransferAuthenticationPush {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
