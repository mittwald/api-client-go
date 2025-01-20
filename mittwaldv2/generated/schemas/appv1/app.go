package appv1

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "actionCapabilities": {"$ref": "#/components/schemas/de.mittwald.v1.app.ActionCapabilities"}
//    "id":
//        type: "string"
//    "name":
//        type: "string"
//    "tags":
//        type: "array"
//        items:
//            type: "string"
//required:
//    - "id"
//    - "name"
//    - "tags"
//description: "An App is to be understood as a manifest for AppInstallations. E.g. 'WordPress' only exists inside our ecosystem, because there is an  App -Manifest for it."

//An App is to be understood as a manifest for AppInstallations. E.g. 'WordPress' only exists inside our ecosystem, because there is an  App -Manifest for it.
type App struct {
	ActionCapabilities ActionCapabilities `json:"actionCapabilities,omitempty"`
	Id                 string             `json:"id"`
	Name               string             `json:"name"`
	Tags               []string           `json:"tags"`
}

func (o *App) Validate() error {
	if err := func() error {
		if o.ActionCapabilities == nil {
			return nil
		}
		return func() error {
			for i := range o.ActionCapabilities {
				if err := o.ActionCapabilities[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property actionCapabilities: %w", err)
	}
	if o.Tags == nil {
		return errors.New("property tags is required, but not set")
	}
	return nil
}
