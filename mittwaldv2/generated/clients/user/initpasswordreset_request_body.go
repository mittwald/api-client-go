package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "email":
//        type: "string"
//        format: "email"
//        description: "Email address to reset the password for."
//        example: "a.lovelace@example.com"
//required:
//    - "email"

//
type InitPasswordResetRequestBody struct {
	Email string `json:"email"`
}

func (o *InitPasswordResetRequestBody) Validate() error {
	return nil
}
