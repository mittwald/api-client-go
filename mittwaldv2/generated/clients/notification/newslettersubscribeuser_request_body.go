package notification

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "firstName":
//        type: "string"
//        example: "Ada"
//    "lastName":
//        type: "string"
//        example: "Lovelace"

//
type NewsletterSubscribeUserRequestBody struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

func (o *NewsletterSubscribeUserRequestBody) Validate() error {
	return nil
}