package user

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/signupv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "sshKeys":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.signup.SshKey"}

//
type ListSSHKeysResponse struct {
	SshKeys []signupv1.SshKey `json:"sshKeys,omitempty"`
}

func (o *ListSSHKeysResponse) Validate() error {
	if err := func() error {
		if o.SshKeys == nil {
			return nil
		}
		return func() error {
			for i := range o.SshKeys {
				if err := o.SshKeys[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property sshKeys: %w", err)
	}
	return nil
}
