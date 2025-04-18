package userclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "multiFactorCode":
//        type: "string"
//        maxLength: 16
//        minLength: 6
//        description: "Multi Factor Code to confirm MFA."
// required:
//    - "multiFactorCode"
// description: DisableMFARequestBody models the JSON body of a 'user-disable-mfa' request

// DisableMFARequestBody models the JSON body of a 'user-disable-mfa' request
type DisableMFARequestBody struct {
	MultiFactorCode string `json:"multiFactorCode"`
}

func (o *DisableMFARequestBody) Validate() error {
	return nil
}
