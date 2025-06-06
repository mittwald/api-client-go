package ingressv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "acme":
//        type: "boolean"
//        description: "Has to be `true`, as ssl cannot be deactivated."
//    "isCreated":
//        type: "boolean"
//    "requestDeadline":
//        type: "string"
//        format: "date-time"
// required:
//    - "acme"
//    - "isCreated"
// additionalProperties: false

type TlsAcme struct {
	Acme            bool       `json:"acme"`
	IsCreated       bool       `json:"isCreated"`
	RequestDeadline *time.Time `json:"requestDeadline,omitempty"`
}

func (o *TlsAcme) Validate() error {
	return nil
}
