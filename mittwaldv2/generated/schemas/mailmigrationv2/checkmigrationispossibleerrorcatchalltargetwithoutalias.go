package mailmigrationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "address":
//        type: "string"
//    "mailboxName":
//        type: "string"
// required:
//    - "address"
//    - "mailboxName"

type CheckMigrationIsPossibleErrorCatchAllTargetWithoutAlias struct {
	Address     string `json:"address"`
	MailboxName string `json:"mailboxName"`
}

func (o *CheckMigrationIsPossibleErrorCatchAllTargetWithoutAlias) Validate() error {
	return nil
}
