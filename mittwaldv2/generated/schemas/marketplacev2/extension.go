package marketplacev2

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "assets":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionAsset"}
//        maxItems: 4
//        description: "The assets/media (images and videos) of the extension."
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
//        description: "A short description of the capabilites of the Extension."
//    "detailedDescriptions": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.DetailedDescriptions"}
//    "disabled":
//        type: "boolean"
//    "externalFrontends":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExternalComponent"}
//    "frontendComponents":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExternalComponent"}
//        deprecated: true
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
//        example: "MyPingExtension"
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
//    "statistics": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionStatistics"}
//    "subTitle": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SubTitle"}
//    "support": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SupportMeta"}
//    "tags":
//        type: "array"
//        items:
//            type: "string"
// required:
//    - "id"
//    - "contributorId"
//    - "support"
//    - "state"
//    - "published"
//    - "name"
//    - "subTitle"
//    - "description"
//    - "tags"
//    - "context"
//    - "scopes"
//    - "disabled"
//    - "blocked"
//    - "assets"
//    - "statistics"
//    - "logoRefId"

type Extension struct {
	Assets               []ExtensionAsset      `json:"assets"`
	Blocked              bool                  `json:"blocked"`
	Context              Context               `json:"context"`
	ContributorId        string                `json:"contributorId"`
	Deprecation          *ExtensionDeprecation `json:"deprecation,omitempty"`
	Description          string                `json:"description"`
	DetailedDescriptions *DetailedDescriptions `json:"detailedDescriptions,omitempty"`
	Disabled             bool                  `json:"disabled"`
	ExternalFrontends    []ExternalComponent   `json:"externalFrontends,omitempty"`
	FrontendComponents   []ExternalComponent   `json:"frontendComponents,omitempty"`
	FrontendFragments    map[string]any        `json:"frontendFragments,omitempty"`
	Id                   string                `json:"id"`
	LogoRefId            string                `json:"logoRefId"`
	Name                 string                `json:"name"`
	Published            bool                  `json:"published"`
	Scopes               []string              `json:"scopes"`
	State                ExtensionState        `json:"state"`
	Statistics           ExtensionStatistics   `json:"statistics"`
	SubTitle             SubTitle              `json:"subTitle"`
	Support              SupportMeta           `json:"support"`
	Tags                 []string              `json:"tags"`
}

func (o *Extension) Validate() error {
	if o.Assets == nil {
		return errors.New("property assets is required, but not set")
	}
	if err := func() error {
		for i := range o.Assets {
			if err := o.Assets[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property assets: %w", err)
	}
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
		if o.ExternalFrontends == nil {
			return nil
		}
		return func() error {
			for i := range o.ExternalFrontends {
				if err := o.ExternalFrontends[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property externalFrontends: %w", err)
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
	if err := o.Statistics.Validate(); err != nil {
		return fmt.Errorf("invalid property statistics: %w", err)
	}
	if err := o.SubTitle.Validate(); err != nil {
		return fmt.Errorf("invalid property subTitle: %w", err)
	}
	if err := o.Support.Validate(); err != nil {
		return fmt.Errorf("invalid property support: %w", err)
	}
	if o.Tags == nil {
		return errors.New("property tags is required, but not set")
	}
	return nil
}
