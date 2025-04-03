package userclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "authorization_code"

type OauthGetAuthorizationRequestQueryGranttype string

const OauthGetAuthorizationRequestQueryGranttypeAuthorizationcode OauthGetAuthorizationRequestQueryGranttype = "authorization_code"

func (e OauthGetAuthorizationRequestQueryGranttype) Validate() error {
	if e == OauthGetAuthorizationRequestQueryGranttypeAuthorizationcode {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
