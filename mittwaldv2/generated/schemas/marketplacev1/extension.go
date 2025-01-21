package marketplacev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "blocked":
//        type: "boolean"
//        deprecated: true
//    "context": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.Context"}
//    "contributorId":
//        type: "string"
//        format: "uuid"
//    "deprecation": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionDeprecation"}
//    "description":
//        type: "string"
//    "detailedDescriptions": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.DetailedDescriptions"}
//    "disabled":
//        type: "boolean"
//    "frontendComponents":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExternalComponent"}
//    "frontendFragments":
//        type: "object"
//        additionalProperties: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.FrontendFragment"}
//    "id":
//        type: "string"
//        format: "uuid"
//    "logoRefId":
//        type: "string"
//        format: "uuid"
//        description: "This is the FileId of the Logo. Retrieve the file with this id on `/v2/files/{logoRefId}`."
//    "name":
//        type: "string"
//    "published":
//        type: "boolean"
//        description: "Whether the extension has been published by the contributor."
//    "scopes":
//        type: "array"
//        items:
//            type: "string"
//    "state":
//        type: "string"
//        enum:
//            - "enabled"
//            - "blocked"
//            - "disabled"
//        description: "deprecated"
//        deprecated: true
//    "support": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SupportMeta"}
//    "tags":
//        type: "array"
//        items:
//            type: "string"
//required:
//    - "id"
//    - "contributorId"
//    - "support"
//    - "state"
//    - "published"
//    - "name"
//    - "description"
//    - "tags"
//    - "context"
//    - "scopes"
//    - "disabled"
//    - "blocked"

type Extension struct {
	Blocked              bool                        `json:"blocked"`
	Context              Context                     `json:"context"`
	ContributorId        uuid.UUID                   `json:"contributorId"`
	Deprecation          *ExtensionDeprecation       `json:"deprecation,omitempty"`
	Description          string                      `json:"description"`
	DetailedDescriptions *DetailedDescriptions       `json:"detailedDescriptions,omitempty"`
	Disabled             bool                        `json:"disabled"`
	FrontendComponents   []ExternalComponent         `json:"frontendComponents,omitempty"`
	FrontendFragments    map[string]FrontendFragment `json:"frontendFragments,omitempty"`
	Id                   uuid.UUID                   `json:"id"`
	LogoRefId            *uuid.UUID                  `json:"logoRefId,omitempty"`
	Name                 string                      `json:"name"`
	Published            bool                        `json:"published"`
	Scopes               []string                    `json:"scopes"`
	State                ExtensionState              `json:"state"`
	Support              SupportMeta                 `json:"support"`
	Tags                 []string                    `json:"tags"`
}

func (o *Extension) Validate() error {
	if err := o.Context.Validate(); err != nil {
		return fmt.Errorf("invalid property context: %w", err)
	}
	if err := func() error {
		if o.Deprecation == nil {
			return nil
		}
		return o.Deprecation.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property deprecation: %w", err)
	}
	if err := func() error {
		if o.DetailedDescriptions == nil {
			return nil
		}
		return o.DetailedDescriptions.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property detailedDescriptions: %w", err)
	}
	if err := func() error {
		if o.FrontendComponents == nil {
			return nil
		}
		return func() error {
			for i := range o.FrontendComponents {
				if err := o.FrontendComponents[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property frontendComponents: %w", err)
	}
	if o.Scopes == nil {
		return errors.New("property scopes is required, but not set")
	}
	if err := o.State.Validate(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	if err := o.Support.Validate(); err != nil {
		return fmt.Errorf("invalid property support: %w", err)
	}
	if o.Tags == nil {
		return errors.New("property tags is required, but not set")
	}
	return nil
}
