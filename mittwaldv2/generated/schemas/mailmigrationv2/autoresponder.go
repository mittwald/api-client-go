package mailmigrationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "active":
//        type: "boolean"
//    "expiresAt":
//        format: "date-time"
//    "message":
//        type: "string"
//    "startsAt":
//        format: "date-time"
// required:
//    - "active"
//    - "message"

type AutoResponder struct {
	Active    bool   `json:"active"`
	ExpiresAt *any   `json:"expiresAt,omitempty"`
	Message   string `json:"message"`
	StartsAt  *any   `json:"startsAt,omitempty"`
}

func (o *AutoResponder) Validate() error {
	return nil
}
