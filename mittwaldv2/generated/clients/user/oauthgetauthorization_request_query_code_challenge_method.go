package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "S256"

type OauthGetAuthorizationRequestQueryCodechallengemethod string

const OauthGetAuthorizationRequestQueryCodechallengemethodS256 OauthGetAuthorizationRequestQueryCodechallengemethod = "S256"

func (e OauthGetAuthorizationRequestQueryCodechallengemethod) Validate() error {
	if e == OauthGetAuthorizationRequestQueryCodechallengemethodS256 {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
