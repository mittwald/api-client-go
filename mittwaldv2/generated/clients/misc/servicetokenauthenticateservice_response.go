package misc

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessToken":
//        type: "object"
//        properties:
//            "id":
//                type: "string"
//            "jwtClaims":
//                type: "object"
//                properties:
//                    "exp":
//                        type: "string"
//                    "iat":
//                        type: "string"
//                    "iss":
//                        type: "string"
//                    "sub":
//                        type: "string"
//                required:
//                    - "iat"
//                    - "iss"
//                    - "sub"
//                    - "exp"
//            "publicToken":
//                type: "string"
//        required:
//            - "id"
//            - "jwtClaims"
//            - "publicToken"
//required:
//    - "accessToken"

//
type ServicetokenAuthenticateServiceResponse struct {
	AccessToken ServicetokenAuthenticateServiceResponseAccessToken `json:"accessToken"`
}

func (o *ServicetokenAuthenticateServiceResponse) Validate() error {
	if err := o.AccessToken.Validate(); err != nil {
		return fmt.Errorf("invalid property accessToken: %w", err)
	}
	return nil
}
