package mailmigrationv2

import "errors"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "address":
//        type: "string"
//    "mailboxes":
//        type: "array"
//        items:
//            type: "string"
// required:
//    - "address"
//    - "mailboxes"

type CheckMigrationIsPossibleErrorAmbiguousMailboxDelivery struct {
	Address   string   `json:"address"`
	Mailboxes []string `json:"mailboxes"`
}

func (o *CheckMigrationIsPossibleErrorAmbiguousMailboxDelivery) Validate() error {
	if o.Mailboxes == nil {
		return errors.New("property mailboxes is required, but not set")
	}
	return nil
}
