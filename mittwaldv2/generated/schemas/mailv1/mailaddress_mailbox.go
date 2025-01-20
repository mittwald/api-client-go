package mailv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "name":
//        type: "string"
//    "passwordUpdatedAt":
//        type: "string"
//        format: "date-time"
//    "sendingEnabled":
//        type: "boolean"
//    "spamProtection":
//        type: "object"
//        properties:
//            "active":
//                type: "boolean"
//            "autoDeleteSpam":
//                type: "boolean"
//            "folder":
//                type: "string"
//                enum:
//                    - "spam"
//                    - "inbox"
//            "relocationMinSpamScore":
//                type: "integer"
//                maximum: 10
//        required:
//            - "active"
//            - "folder"
//            - "relocationMinSpamScore"
//            - "autoDeleteSpam"
//    "storageInBytes":
//        type: "object"
//        properties:
//            "current":
//                type: "object"
//                properties:
//                    "updatedAt":
//                        type: "string"
//                        format: "date-time"
//                    "value":
//                        type: "number"
//                required:
//                    - "value"
//                    - "updatedAt"
//            "limit":
//                type: "number"
//        required:
//            - "limit"
//            - "current"
//required:
//    - "name"
//    - "sendingEnabled"
//    - "spamProtection"
//    - "storageInBytes"
//    - "passwordUpdatedAt"

//
type MailAddressMailbox struct {
	Name              string                           `json:"name"`
	PasswordUpdatedAt string                           `json:"passwordUpdatedAt"`
	SendingEnabled    bool                             `json:"sendingEnabled"`
	SpamProtection    MailAddressMailboxSpamProtection `json:"spamProtection"`
	StorageInBytes    MailAddressMailboxStorageInBytes `json:"storageInBytes"`
}

func (o *MailAddressMailbox) Validate() error {
	if err := o.SpamProtection.Validate(); err != nil {
		return fmt.Errorf("invalid property spamProtection: %w", err)
	}
	if err := o.StorageInBytes.Validate(); err != nil {
		return fmt.Errorf("invalid property storageInBytes: %w", err)
	}
	return nil
}
