package mailv1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "updatedAt":
//        type: "string"
//        format: "date-time"
//    "value":
//        type: "number"
//required:
//    - "value"
//    - "updatedAt"

//
type MailAddressMailboxStorageInBytesCurrent struct {
	UpdatedAt time.Time `json:"updatedAt"`
	Value     float64   `json:"value"`
}

func (o *MailAddressMailboxStorageInBytesCurrent) Validate() error {
	return nil
}
