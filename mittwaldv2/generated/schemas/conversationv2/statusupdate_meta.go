package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "user": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}

type StatusUpdateMeta struct {
	User *User `json:"user,omitempty"`
}

func (o *StatusUpdateMeta) Validate() error {
	if err := func() error {
		if o.User == nil {
			return nil
		}
		return o.User.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property user: %w", err)
	}
	return nil
}
