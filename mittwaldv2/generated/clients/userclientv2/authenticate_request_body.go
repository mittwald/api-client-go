package userclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "email":
//        type: "string"
//        description: "The users Email-Address."
//        example: "a.lovelace@example.com"
//    "password":
//        type: "string"
//        description: "Password of the User."
// required:
//    - "email"
//    - "password"
// description: AuthenticateRequestBody models the JSON body of a 'user-authenticate' request

// AuthenticateRequestBody models the JSON body of a 'user-authenticate' request
type AuthenticateRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (o *AuthenticateRequestBody) Validate() error {
	return nil
}
