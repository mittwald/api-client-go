package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "token":
//        type: "string"
//        description: "The `ApiToken`."
//required:
//    - "token"

//
type CreateAPITokenResponse struct {
	Token string `json:"token"`
}

func (o *CreateAPITokenResponse) Validate() error {
	return nil
}
