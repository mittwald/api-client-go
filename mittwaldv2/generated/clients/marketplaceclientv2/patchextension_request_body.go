package marketplaceclientv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "deprecation": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExtensionDeprecation"}
//    "description":
//        type: "string"
//    "detailedDescriptions": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.DetailedDescriptions"}
//    "externalFrontends":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.ExternalComponent"}
//    "frontendFragments":
//        type: "object"
//        additionalProperties: {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.FrontendFragment"}
//    "name":
//        type: "string"
//    "scopes":
//        type: "array"
//        items:
//            type: "string"
//    "subTitle": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SubTitle"}
//    "support": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.SupportMeta"}
//    "tags":
//        type: "array"
//        items:
//            type: "string"
//    "webhookUrls": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.WebhookUrls"}
// description: PatchExtensionRequestBody models the JSON body of a 'extension-patch-extension' request

// PatchExtensionRequestBody models the JSON body of a 'extension-patch-extension' request
type PatchExtensionRequestBody struct {
	Deprecation          *marketplacev2.ExtensionDeprecation `json:"deprecation,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	DetailedDescriptions *marketplacev2.DetailedDescriptions `json:"detailedDescriptions,omitempty"`
	ExternalFrontends    []marketplacev2.ExternalComponent   `json:"externalFrontends,omitempty"`
	FrontendFragments    map[string]any                      `json:"frontendFragments,omitempty"`
	Name                 *string                             `json:"name,omitempty"`
	Scopes               []string                            `json:"scopes,omitempty"`
	SubTitle             *marketplacev2.SubTitle             `json:"subTitle,omitempty"`
	Support              *marketplacev2.SupportMeta          `json:"support,omitempty"`
	Tags                 []string                            `json:"tags,omitempty"`
	WebhookUrls          *marketplacev2.WebhookUrls          `json:"webhookUrls,omitempty"`
}

func (o *PatchExtensionRequestBody) Validate() error {
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
		if o.Scopes == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property scopes: %w", err)
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
