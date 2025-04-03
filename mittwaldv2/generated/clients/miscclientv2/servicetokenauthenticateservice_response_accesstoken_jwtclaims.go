package miscclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "exp":
//        type: "string"
//    "iat":
//        type: "string"
//    "iss":
//        type: "string"
//    "sub":
//        type: "string"
// required:
//    - "iat"
//    - "iss"
//    - "sub"
//    - "exp"

type ServicetokenAuthenticateServiceResponseAccessTokenJwtClaims struct {
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
}

func (o *ServicetokenAuthenticateServiceResponseAccessTokenJwtClaims) Validate() error {
	return nil
}
