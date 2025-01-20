package user

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "comment":
//        type: "string"
//        example: "a.lovelace@example.com"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "comment"

//
type EditSSHKeyRequestBody struct {
	Comment   string     `json:"comment"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

func (o *EditSSHKeyRequestBody) Validate() error {
	return nil
}