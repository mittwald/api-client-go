package sshuserv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessLevel": {"$ref": "#/components/schemas/de.mittwald.v1.sshuser.AccessLevel"}
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
//    "directories":
//        type: "array"
//        items:
//            type: "string"
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
//    - "accessLevel"
//    - "hasPassword"
//description: "A representation of an SFTPUser."

//A representation of an SFTPUser.
type SftpUser struct {
	AccessLevel   AccessLevel `json:"accessLevel"`
	Active        *bool       `json:"active,omitempty"`
	AuthUpdatedAt string      `json:"authUpdatedAt"`
	CreatedAt     string      `json:"createdAt"`
	Description   string      `json:"description"`
	Directories   []string    `json:"directories,omitempty"`
	ExpiresAt     *string     `json:"expiresAt,omitempty"`
	HasPassword   bool        `json:"hasPassword"`
	Id            string      `json:"id"`
	ProjectId     string      `json:"projectId"`
	PublicKeys    []PublicKey `json:"publicKeys,omitempty"`
	UpdatedAt     *string     `json:"updatedAt,omitempty"`
	UserName      string      `json:"userName"`
}

func (o *SftpUser) Validate() error {
	if err := o.AccessLevel.Validate(); err != nil {
		return fmt.Errorf("invalid property accessLevel: %w", err)
	}
	if err := func() error {
		if o.Directories == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property directories: %w", err)
	}
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
