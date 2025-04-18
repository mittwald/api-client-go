package userclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "SecondFactorRequired"

type AuthenticateAcceptedResponseName string

const AuthenticateAcceptedResponseNameSecondFactorRequired AuthenticateAcceptedResponseName = "SecondFactorRequired"

func (e AuthenticateAcceptedResponseName) Validate() error {
	if e == AuthenticateAcceptedResponseNameSecondFactorRequired {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
