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
//    "backendComponents": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.BackendComponents"}
//    "blocked":
//        type: "boolean"
//        deprecated: true
//    "context": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.Context"}
//    "contributorId":
//        type: "string"
//    "deprecation": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionDeprecation"}
//    "description":
//        type: "string"
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
//    "functional":
//        type: "boolean"
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
//    "requestedChanges":
//        type: "object"
//        properties:
//            "context": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.Context"}
//            "scopes":
//                type: "array"
//                items:
//                    type: "string"
//            "webhookUrls": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.WebhookUrls"}
//        additionalProperties: false
//    "scopes":
//        type: "array"
//        items:
//            type: "string"
//    "secrets":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionSecret"}
//    "state":
//        type: "string"
//        enum:
//            - "enabled"
//            - "blocked"
//            - "disabled"
//        description: "deprecated"
//    "statistics": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionStatistics"}
//    "subTitle": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SubTitle"}
//    "support": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SupportMeta"}
//    "tags":
//        type: "array"
//        items:
//            type: "string"
//    "verificationRequested":
//        type: "boolean"
//    "verified":
//        type: "boolean"
//    "webhookUrls": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.WebhookUrls"}
// required:
//    - "id"
//    - "contributorId"
//    - "name"
//    - "statistics"
//    - "assets"
//    - "published"
//    - "verified"
//    - "verificationRequested"
//    - "functional"
//    - "secrets"

type OwnExtension struct {
	Assets                []ExtensionAsset              `json:"assets"`
	BackendComponents     *BackendComponents            `json:"backendComponents,omitempty"`
	Blocked               *bool                         `json:"blocked,omitempty"`
	Context               *Context                      `json:"context,omitempty"`
	ContributorId         string                        `json:"contributorId"`
	Deprecation           *ExtensionDeprecation         `json:"deprecation,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	DetailedDescriptions  *DetailedDescriptions         `json:"detailedDescriptions,omitempty"`
	Disabled              *bool                         `json:"disabled,omitempty"`
	ExternalFrontends     []ExternalComponent           `json:"externalFrontends,omitempty"`
	FrontendComponents    []ExternalComponent           `json:"frontendComponents,omitempty"`
	FrontendFragments     map[string]any                `json:"frontendFragments,omitempty"`
	Functional            bool                          `json:"functional"`
	Id                    string                        `json:"id"`
	LogoRefId             *string                       `json:"logoRefId,omitempty"`
	Name                  string                        `json:"name"`
	Published             bool                          `json:"published"`
	RequestedChanges      *OwnExtensionRequestedChanges `json:"requestedChanges,omitempty"`
	Scopes                []string                      `json:"scopes,omitempty"`
	Secrets               []ExtensionSecret             `json:"secrets"`
	State                 *OwnExtensionState            `json:"state,omitempty"`
	Statistics            ExtensionStatistics           `json:"statistics"`
	SubTitle              *SubTitle                     `json:"subTitle,omitempty"`
	Support               *SupportMeta                  `json:"support,omitempty"`
	Tags                  []string                      `json:"tags,omitempty"`
	VerificationRequested bool                          `json:"verificationRequested"`
	Verified              bool                          `json:"verified"`
	WebhookUrls           *WebhookUrls                  `json:"webhookUrls,omitempty"`
}

func (o *OwnExtension) Validate() error {
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
	if err := func() error {
		if o.BackendComponents == nil {
			return nil
		}
		return o.BackendComponents.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property backendComponents: %w", err)
	}
	if err := func() error {
		if o.Context == nil {
			return nil
		}
		return o.Context.Validate()
	}(); err != nil {
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
	if err := func() error {
		if o.RequestedChanges == nil {
			return nil
		}
		return o.RequestedChanges.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property requestedChanges: %w", err)
	}
	if err := func() error {
		if o.Scopes == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property scopes: %w", err)
	}
	if o.Secrets == nil {
		return errors.New("property secrets is required, but not set")
	}
	if err := func() error {
		for i := range o.Secrets {
			if err := o.Secrets[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property secrets: %w", err)
	}
	if err := func() error {
		if o.State == nil {
			return nil
		}
		return o.State.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	if err := o.Statistics.Validate(); err != nil {
		return fmt.Errorf("invalid property statistics: %w", err)
	}
	if err := func() error {
		if o.SubTitle == nil {
			return nil
		}
		return o.SubTitle.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property subTitle: %w", err)
	}
	if err := func() error {
		if o.Support == nil {
			return nil
		}
		return o.Support.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property support: %w", err)
	}
	if err := func() error {
		if o.Tags == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property tags: %w", err)
	}
	if err := func() error {
		if o.WebhookUrls == nil {
			return nil
		}
		return o.WebhookUrls.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property webhookUrls: %w", err)
	}
	return nil
}
