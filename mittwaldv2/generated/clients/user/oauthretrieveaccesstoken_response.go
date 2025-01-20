package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "access_token":
//        type: "string"
//        description: "The access token issued by the authorization server.\n"
//    "expires_in":
//        type: "integer"
//        description: "The lifetime in seconds of the access token. For\nexample, the value \"3600\" denotes that the access\ntoken will expire in one hour from the time the\nresponse was generated.\n"
//    "refresh_token":
//        type: "string"
//        description: "The refresh token issued by the authorization server.\n"
//    "scope":
//        type: "string"
//        description: "The scope of the access token as described by\n[RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-3.3).\n"
//    "token_type":
//        type: "string"
//        enum:
//            - "bearer"
//        description: "The type of the token issued as described in\n[RFC6749](https://datatracker.ietf.org/doc/html/rfc6749#section-7.1).\n"
//required:
//    - "access_token"
//    - "refresh_token"
//    - "token_type"
//    - "expires_in"

//
type OauthRetrieveAccessTokenResponse struct {
	Access_token  string                                    `json:"access_token"`
	Expires_in    int64                                     `json:"expires_in"`
	Refresh_token string                                    `json:"refresh_token"`
	Scope         *string                                   `json:"scope,omitempty"`
	Token_type    OauthRetrieveAccessTokenResponseTokentype `json:"token_type"`
}

func (o *OauthRetrieveAccessTokenResponse) Validate() error {
	if err := o.Token_type.Validate(); err != nil {
		return fmt.Errorf("invalid property token_type: %w", err)
	}
	return nil
}