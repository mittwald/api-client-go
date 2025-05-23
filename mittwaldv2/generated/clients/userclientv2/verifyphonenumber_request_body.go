package userclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "code":
//        type: "string"
//        maxLength: 6
//        minLength: 6
//        example: "123456"
//    "phoneNumber":
//        type: "string"
//        example: "+491701234567"
// required:
//    - "phoneNumber"
//    - "code"
// description: VerifyPhoneNumberRequestBody models the JSON body of a 'user-verify-phone-number' request

// VerifyPhoneNumberRequestBody models the JSON body of a 'user-verify-phone-number' request
type VerifyPhoneNumberRequestBody struct {
	Code        string `json:"code"`
	PhoneNumber string `json:"phoneNumber"`
}

func (o *VerifyPhoneNumberRequestBody) Validate() error {
	return nil
}
