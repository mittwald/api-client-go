package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "active":
//        type: "boolean"
//    "avatarRefId":
//        type: "string"
//    "clearName":
//        type: "string"
//    "department": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.Department"}
//    "userId":
//        type: "string"
// required:
//    - "userId"

type User struct {
	Active      *bool       `json:"active,omitempty"`
	AvatarRefId *string     `json:"avatarRefId,omitempty"`
	ClearName   *string     `json:"clearName,omitempty"`
	Department  *Department `json:"department,omitempty"`
	UserId      string      `json:"userId"`
}

func (o *User) Validate() error {
	if err := func() error {
		if o.Department == nil {
			return nil
		}
		return o.Department.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property department: %w", err)
	}
	return nil
}
