package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessTokenRetrievalKey":
//        type: "string"
//        maxLength: 37
//        minLength: 37
//    "userId":
//        type: "string"
//        format: "uuid"
//required:
//    - "userId"
//    - "accessTokenRetrievalKey"
//    - "refreshToken"

type AuthenticateWithAccessTokenRetrievalKeyRequestBody struct {
	AccessTokenRetrievalKey string    `json:"accessTokenRetrievalKey"`
	UserId                  uuid.UUID `json:"userId"`
}

func (o *AuthenticateWithAccessTokenRetrievalKeyRequestBody) Validate() error {
	return nil
}
