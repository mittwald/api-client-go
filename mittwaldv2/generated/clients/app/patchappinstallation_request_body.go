package app

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/appv1"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "appVersionId":
//        type: "string"
//        format: "uuid"
//    "customDocumentRoot":
//        type: "string"
//    "description":
//        type: "string"
//    "systemSoftware":
//        type: "object"
//        additionalProperties:
//            type: "object"
//            properties:
//                "systemSoftwareVersion":
//                    type: "string"
//                "updatePolicy": {"$ref": "#/components/schemas/de.mittwald.v1.app.SystemSoftwareUpdatePolicy"}
//    "updatePolicy": {"$ref": "#/components/schemas/de.mittwald.v1.app.AppUpdatePolicy"}
//    "userInputs":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.SavedUserInput"}

//
type PatchAppinstallationRequestBody struct {
	AppVersionId       *uuid.UUID                                                   `json:"appVersionId,omitempty"`
	CustomDocumentRoot *string                                                      `json:"customDocumentRoot,omitempty"`
	Description        *string                                                      `json:"description,omitempty"`
	SystemSoftware     map[string]PatchAppinstallationRequestBodySystemSoftwareItem `json:"systemSoftware,omitempty"`
	UpdatePolicy       *appv1.AppUpdatePolicy                                       `json:"updatePolicy,omitempty"`
	UserInputs         []appv1.SavedUserInput                                       `json:"userInputs,omitempty"`
}

func (o *PatchAppinstallationRequestBody) Validate() error {
	if err := func() error {
		if o.UpdatePolicy == nil {
			return nil
		}
		return o.UpdatePolicy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property updatePolicy: %w", err)
	}
	if err := func() error {
		if o.UserInputs == nil {
			return nil
		}
		return func() error {
			for i := range o.UserInputs {
				if err := o.UserInputs[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property userInputs: %w", err)
	}
	return nil
}
