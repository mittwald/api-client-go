package app

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/appv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "systemSoftwareVersion":
//        type: "string"
//    "updatePolicy": {"$ref": "#/components/schemas/de.mittwald.v1.app.SystemSoftwareUpdatePolicy"}

//
type PatchAppinstallationRequestBodySystemSoftwareItem struct {
	SystemSoftwareVersion *string                           `json:"systemSoftwareVersion,omitempty"`
	UpdatePolicy          *appv1.SystemSoftwareUpdatePolicy `json:"updatePolicy,omitempty"`
}

func (o *PatchAppinstallationRequestBodySystemSoftwareItem) Validate() error {
	if err := func() error {
		if o.UpdatePolicy == nil {
			return nil
		}
		return o.UpdatePolicy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property updatePolicy: %w", err)
	}
	return nil
}