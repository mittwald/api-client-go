package sshuserv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "password":
//        type: "string"
//        maxLength: 72
// required:
//    - "password"

type AuthenticationAlternative1 struct {
	Password string `json:"password"`
}

func (o *AuthenticationAlternative1) Validate() error {
	return nil
}
