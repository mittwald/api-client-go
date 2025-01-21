package user

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "description":
//        type: "string"
//        example: "Api Token for ..."
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "description"

type EditAPITokenRequestBody struct {
	Description string     `json:"description"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
}

func (o *EditAPITokenRequestBody) Validate() error {
	return nil
}
