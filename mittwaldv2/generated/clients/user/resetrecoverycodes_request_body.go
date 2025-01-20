package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "multiFactorCode":
//        type: "string"
//        maxLength: 16
//        minLength: 6
//        description: "Multi Factor Code to authorize your request."
//        example: "123456"
//required:
//    - "multiFactorCode"

//
type ResetRecoverycodesRequestBody struct {
	MultiFactorCode string `json:"multiFactorCode"`
}

func (o *ResetRecoverycodesRequestBody) Validate() error {
	return nil
}
