package domainclientv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "acme":
//        type: "boolean"
//    "isCreated":
//        type: "boolean"
//        description: "Was added by mistake. Never did anything."
//        deprecated: true
//    "requestDeadline":
//        type: "string"
//        format: "date-time"
//        description: "Was added by mistake. Never did anything."
//        deprecated: true
// required:
//    - "acme"
// additionalProperties: false

type UpdateIngressTLSRequestBodyAlternative1 struct {
	Acme            bool       `json:"acme"`
	IsCreated       *bool      `json:"isCreated,omitempty"`
	RequestDeadline *time.Time `json:"requestDeadline,omitempty"`
}

func (o *UpdateIngressTLSRequestBodyAlternative1) Validate() error {
	return nil
}
