package sshuserv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "active":
//        type: "boolean"
//    "authUpdatedAt":
//        type: "string"
//        format: "date-time"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "description":
//        type: "string"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//    "hasPassword":
//        type: "boolean"
//    "id":
//        type: "string"
//    "projectId":
//        type: "string"
//    "publicKeys":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.sshuser.PublicKey"}
//    "updatedAt":
//        type: "string"
//        format: "date-time"
//    "userName":
//        type: "string"
//required:
//    - "id"
//    - "projectId"
//    - "description"
//    - "userName"
//    - "createdAt"
//    - "authUpdatedAt"
//    - "hasPassword"
//description: "A representation of an SSHUser."

//A representation of an SSHUser.
type SshUser struct {
	Active        *bool       `json:"active,omitempty"`
	AuthUpdatedAt time.Time   `json:"authUpdatedAt"`
	CreatedAt     time.Time   `json:"createdAt"`
	Description   string      `json:"description"`
	ExpiresAt     *time.Time  `json:"expiresAt,omitempty"`
	HasPassword   bool        `json:"hasPassword"`
	Id            string      `json:"id"`
	ProjectId     string      `json:"projectId"`
	PublicKeys    []PublicKey `json:"publicKeys,omitempty"`
	UpdatedAt     *time.Time  `json:"updatedAt,omitempty"`
	UserName      string      `json:"userName"`
}

func (o *SshUser) Validate() error {
	if err := func() error {
		if o.PublicKeys == nil {
			return nil
		}
		return func() error {
			for i := range o.PublicKeys {
				if err := o.PublicKeys[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property publicKeys: %w", err)
	}
	return nil
}
