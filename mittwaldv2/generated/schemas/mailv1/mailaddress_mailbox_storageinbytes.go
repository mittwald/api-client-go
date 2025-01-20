package mailv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "current":
//        type: "object"
//        properties:
//            "updatedAt":
//                type: "string"
//                format: "date-time"
//            "value":
//                type: "number"
//        required:
//            - "value"
//            - "updatedAt"
//    "limit":
//        type: "number"
//required:
//    - "limit"
//    - "current"

//
type MailAddressMailboxStorageInBytes struct {
	Current MailAddressMailboxStorageInBytesCurrent `json:"current"`
	Limit   float64                                 `json:"limit"`
}

func (o *MailAddressMailboxStorageInBytes) Validate() error {
	if err := o.Current.Validate(); err != nil {
		return fmt.Errorf("invalid property current: %w", err)
	}
	return nil
}
