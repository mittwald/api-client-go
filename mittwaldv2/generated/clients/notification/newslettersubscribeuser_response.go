package notification

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "active":
//        type: "boolean"
//        example: false
//    "email":
//        type: "string"
//        example: "a.lovelace@example.com"
//    "registered":
//        type: "boolean"
//        example: true
//required:
//    - "email"
//    - "active"
//    - "registered"

type NewsletterSubscribeUserResponse struct {
	Active     bool   `json:"active"`
	Email      string `json:"email"`
	Registered bool   `json:"registered"`
}

func (o *NewsletterSubscribeUserResponse) Validate() error {
	return nil
}
