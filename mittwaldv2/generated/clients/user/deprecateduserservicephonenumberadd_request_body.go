package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "phoneNumber":
//        type: "string"
//required:
//    - "phoneNumber"

//
type DeprecatedUserServicePhoneNumberAddRequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

func (o *DeprecatedUserServicePhoneNumberAddRequestBody) Validate() error {
	return nil
}
