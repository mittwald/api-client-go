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
//    "authentication": {"$ref": "#/components/schemas/de.mittwald.v1.sshuser.Authentication"}
//    "description":
//        type: "string"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "description"
//    - "authentication"

//
type CreateSSHUserRequestBody struct {
	Authentication sshuserv1.Authentication `json:"authentication"`
	Description    string                   `json:"description"`
	ExpiresAt      *time.Time               `json:"expiresAt,omitempty"`
}

func (o *CreateSSHUserRequestBody) Validate() error {
	if err := o.Authentication.Validate(); err != nil {
		return fmt.Errorf("invalid property authentication: %w", err)
	}
	return nil
}
