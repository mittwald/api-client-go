package mailclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "autoResponder":
//        type: "object"
//        properties:
//            "active":
//                type: "boolean"
//            "expiresAt":
//                type: "string"
//                format: "date-time"
//            "message":
//                type: "string"
//            "startsAt":
//                type: "string"
//                format: "date-time"
//        required:
//            - "message"
//            - "active"
//        nullable: true
// required:
//    - "autoResponder"
// description: UpdateMailAddressAutoresponderRequestBody models the JSON body of a 'mail-update-mail-address-autoresponder' request

// UpdateMailAddressAutoresponderRequestBody models the JSON body of a 'mail-update-mail-address-autoresponder' request
type UpdateMailAddressAutoresponderRequestBody struct {
	AutoResponder UpdateMailAddressAutoresponderRequestBodyAutoResponder `json:"autoResponder"`
}

func (o *UpdateMailAddressAutoresponderRequestBody) Validate() error {
	if err := o.AutoResponder.Validate(); err != nil {
		return fmt.Errorf("invalid property autoResponder: %w", err)
	}
	return nil
}
