package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "password":
//        type: "string"
//        description: "The new password."
//    "token":
//        type: "string"
//        maxLength: 6
//        minLength: 6
//        description: "Password reset token"
//    "userId":
//        type: "string"
//        format: "uuid"
//        description: "UserId of the user to reset the password for."
//required:
//    - "userId"
//    - "token"
//    - "password"

//
type ConfirmPasswordResetRequestBody struct {
	Password string    `json:"password"`
	Token    string    `json:"token"`
	UserId   uuid.UUID `json:"userId"`
}

func (o *ConfirmPasswordResetRequestBody) Validate() error {
	return nil
}