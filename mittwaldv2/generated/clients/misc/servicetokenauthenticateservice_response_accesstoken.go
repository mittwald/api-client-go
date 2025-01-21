package misc

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "id":
//        type: "string"
//    "jwtClaims":
//        type: "object"
//        properties:
//            "exp":
//                type: "string"
//            "iat":
//                type: "string"
//            "iss":
//                type: "string"
//            "sub":
//                type: "string"
//        required:
//            - "iat"
//            - "iss"
//            - "sub"
//            - "exp"
//    "publicToken":
//        type: "string"
//required:
//    - "id"
//    - "jwtClaims"
//    - "publicToken"

type ServicetokenAuthenticateServiceResponseAccessToken struct {
	Id          string                                                      `json:"id"`
	JwtClaims   ServicetokenAuthenticateServiceResponseAccessTokenJwtClaims `json:"jwtClaims"`
	PublicToken string                                                      `json:"publicToken"`
}

func (o *ServicetokenAuthenticateServiceResponseAccessToken) Validate() error {
	if err := o.JwtClaims.Validate(); err != nil {
		return fmt.Errorf("invalid property jwtClaims: %w", err)
	}
	return nil
}
