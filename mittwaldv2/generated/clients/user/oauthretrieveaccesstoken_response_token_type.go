package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "bearer"
//description: "The type of the token issued as described in\n[RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-7.1).\n"

//The type of the token issued as described in
//[RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-7.1).
//
type OauthRetrieveAccessTokenResponseTokentype string

const OauthRetrieveAccessTokenResponseTokentypeBearer OauthRetrieveAccessTokenResponseTokentype = "bearer"

func (e OauthRetrieveAccessTokenResponseTokentype) Validate() error {
	if e == OauthRetrieveAccessTokenResponseTokentypeBearer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}