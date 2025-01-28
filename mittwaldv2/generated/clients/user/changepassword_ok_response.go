package user

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "expires":
//        type: "string"
//        format: "date-time"
//        description: "The expiration date of the token."
//    "token":
//        type: "string"
//        description: "Public token to identify yourself against the api gateway."
// required:
//    - "token"
//    - "expires"

type ChangePasswordOKResponse struct {
	Expires time.Time `json:"expires"`
	Token   string    `json:"token"`
}

func (o *ChangePasswordOKResponse) Validate() error {
	return nil
}
