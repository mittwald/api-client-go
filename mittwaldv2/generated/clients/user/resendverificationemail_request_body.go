package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "email":
//        type: "string"
//        format: "email"
//        example: "ada.lovelace@example.com"
//    "userId":
//        type: "string"
//        format: "uuid"
//required:
//    - "userId"
//    - "email"

type ResendVerificationEmailRequestBody struct {
	Email  string    `json:"email"`
	UserId uuid.UUID `json:"userId"`
}

func (o *ResendVerificationEmailRequestBody) Validate() error {
	return nil
}
