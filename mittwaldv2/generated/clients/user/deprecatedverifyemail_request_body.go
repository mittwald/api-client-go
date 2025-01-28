package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "email":
//        type: "string"
//        format: "email"
//        description: "The Email-Address to verify."
//        example: "ada.lovelace@example.com"
//    "token":
//        type: "string"
//        description: "The token found in the verification email."
//        example: "123456"
// required:
//    - "email"
// description: DeprecatedVerifyEmailRequestBody models the JSON body of a 'deprecated-user-verify-email' request

// DeprecatedVerifyEmailRequestBody models the JSON body of a 'deprecated-user-verify-email' request
type DeprecatedVerifyEmailRequestBody struct {
	Email string  `json:"email"`
	Token *string `json:"token,omitempty"`
}

func (o *DeprecatedVerifyEmailRequestBody) Validate() error {
	return nil
}
