package mailv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "active":
//        type: "boolean"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//    "message":
//        type: "string"
//    "startsAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "active"

//
type MailAddressAutoResponder struct {
	Active    bool    `json:"active"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	Message   *string `json:"message,omitempty"`
	StartsAt  *string `json:"startsAt,omitempty"`
}

func (o *MailAddressAutoResponder) Validate() error {
	return nil
}
