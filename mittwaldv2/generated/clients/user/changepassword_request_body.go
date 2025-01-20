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
//        description: "Multi Factor Code to confirm MFA.\nThis is optional, depending on the MFA activation status of the profile.\n"
//        example: "123456"
//    "newPassword":
//        type: "string"
//        description: "The new password."
//    "oldPassword":
//        type: "string"
//        description: "The old password."
//required:
//    - "oldPassword"
//    - "newPassword"

//
type ChangePasswordRequestBody struct {
	MultiFactorCode *string `json:"multiFactorCode,omitempty"`
	NewPassword     string  `json:"newPassword"`
	OldPassword     string  `json:"oldPassword"`
}

func (o *ChangePasswordRequestBody) Validate() error {
	return nil
}