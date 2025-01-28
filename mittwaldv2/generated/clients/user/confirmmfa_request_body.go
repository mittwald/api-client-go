package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "multiFactorCode":
//        type: "string"
//        maxLength: 16
//        minLength: 6
//        description: "Multi Factor Code to confirm MFA."
//        example: "123456"
// required:
//    - "multiFactorCode"
// description: ConfirmMFARequestBody models the JSON body of a 'user-confirm-mfa' request

// ConfirmMFARequestBody models the JSON body of a 'user-confirm-mfa' request
type ConfirmMFARequestBody struct {
	MultiFactorCode string `json:"multiFactorCode"`
}

func (o *ConfirmMFARequestBody) Validate() error {
	return nil
}
