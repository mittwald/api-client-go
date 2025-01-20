package sshsftpuser

import (
	"fmt"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sshuserv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "active":
//        type: "boolean"
//    "description":
//        type: "string"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//    "password":
//        type: "string"
//        maxLength: 72
//    "publicKeys":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.sshuser.PublicKey"}

//
type UpdateSSHUserRequestBody struct {
	Active      *bool                 `json:"active,omitempty"`
	Description *string               `json:"description,omitempty"`
	ExpiresAt   *time.Time            `json:"expiresAt,omitempty"`
	Password    *string               `json:"password,omitempty"`
	PublicKeys  []sshuserv1.PublicKey `json:"publicKeys,omitempty"`
}

func (o *UpdateSSHUserRequestBody) Validate() error {
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
