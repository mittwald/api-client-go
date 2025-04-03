package userclientv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/signupv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "sshKey": {"$ref": "#/components/schemas/de.mittwald.v1.signup.SshKey"}
// required:
//    - "sshKey"

type GetSSHKeyResponse struct {
	SshKey signupv2.SshKey `json:"sshKey"`
}

func (o *GetSSHKeyResponse) Validate() error {
	if err := o.SshKey.Validate(); err != nil {
		return fmt.Errorf("invalid property sshKey: %w", err)
	}
	return nil
}
