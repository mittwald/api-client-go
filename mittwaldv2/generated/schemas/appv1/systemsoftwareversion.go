package appv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/feev1"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "expiryDate":
//        type: "string"
//        format: "date-time"
//    "externalVersion":
//        type: "string"
//    "fee": {"$ref": "#/components/schemas/de.mittwald.v1.fee.FeeStrategy"}
//    "id":
//        type: "string"
//        format: "uuid"
//    "internalVersion":
//        type: "string"
//    "recommended":
//        type: "boolean"
//    "systemSoftwareDependencies":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.SystemSoftwareDependency"}
//    "userInputs":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.UserInput"}
//required:
//    - "id"
//    - "internalVersion"
//    - "externalVersion"
//description: "A SystemSoftwareVersion is an officially  supported version of a SystemSoftware, containing the necessary and recommended configuration und dependencies."

// A SystemSoftwareVersion is an officially  supported version of a SystemSoftware, containing the necessary and recommended configuration und dependencies.
type SystemSoftwareVersion struct {
	ExpiryDate                 *time.Time                 `json:"expiryDate,omitempty"`
	ExternalVersion            string                     `json:"externalVersion"`
	Fee                        *feev1.FeeStrategy         `json:"fee,omitempty"`
	Id                         uuid.UUID                  `json:"id"`
	InternalVersion            string                     `json:"internalVersion"`
	Recommended                *bool                      `json:"recommended,omitempty"`
	SystemSoftwareDependencies []SystemSoftwareDependency `json:"systemSoftwareDependencies,omitempty"`
	UserInputs                 []UserInput                `json:"userInputs,omitempty"`
}

func (o *SystemSoftwareVersion) Validate() error {
	if err := func() error {
		if o.Fee == nil {
			return nil
		}
		return o.Fee.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property fee: %w", err)
	}
	if err := func() error {
		if o.SystemSoftwareDependencies == nil {
			return nil
		}
		return func() error {
			for i := range o.SystemSoftwareDependencies {
				if err := o.SystemSoftwareDependencies[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property systemSoftwareDependencies: %w", err)
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
