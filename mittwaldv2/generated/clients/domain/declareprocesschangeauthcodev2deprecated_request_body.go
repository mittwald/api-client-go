package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "authCode":
//        type: "string"
//        minLength: 4
//required:
//    - "authCode"

//
type DeclareProcessChangeAuthcodeV2DeprecatedRequestBody struct {
	AuthCode string `json:"authCode"`
}

func (o *DeclareProcessChangeAuthcodeV2DeprecatedRequestBody) Validate() error {
	return nil
}