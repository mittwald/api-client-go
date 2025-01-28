package mailmigrationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "address":
//        type: "string"
//    "autoResponder": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.AutoResponder"}
//    "forwardAddresses":
//        type: "array"
//        items:
//            type: "string"
//    "isCatchAll":
//        type: "boolean"
//    "mailbox": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.Mailbox"}
//    "projectId":
//        type: "string"
// required:
//    - "projectId"

type MigrateMailAddressCommandRequirements struct {
	Address          *string        `json:"address,omitempty"`
	AutoResponder    *AutoResponder `json:"autoResponder,omitempty"`
	ForwardAddresses []string       `json:"forwardAddresses,omitempty"`
	IsCatchAll       *bool          `json:"isCatchAll,omitempty"`
	Mailbox          *Mailbox       `json:"mailbox,omitempty"`
	ProjectId        string         `json:"projectId"`
}

func (o *MigrateMailAddressCommandRequirements) Validate() error {
	if err := func() error {
		if o.AutoResponder == nil {
			return nil
		}
		return o.AutoResponder.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property autoResponder: %w", err)
	}
	if err := func() error {
		if o.ForwardAddresses == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property forwardAddresses: %w", err)
	}
	if err := func() error {
		if o.Mailbox == nil {
			return nil
		}
		return o.Mailbox.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property mailbox: %w", err)
	}
	return nil
}
