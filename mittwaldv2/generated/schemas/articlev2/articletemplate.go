package articlev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "additionalArticles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.article.ReadableBookableArticleOptions"}
//    "addons":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.article.ArticleAddons"}
//    "attributes":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.article.ArticleAttributes"}
//    "description":
//        type: "string"
//        example: "Space-Storage"
//    "id":
//        type: "string"
//        minLength: 1
//        example: "054e27e4-d3ed-4ffc-a472-fbb74a6a2ec1"
//    "isManagedByDomain":
//        type: "boolean"
//    "isRecurring":
//        type: "boolean"
//    "modifierArticles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.article.ReadableModifierArticleOptions"}
//    "name":
//        type: "string"
//        minLength: 3
//        example: "Speicher für Hosting aller Art"
//    "type":
//        type: "string"
//        enum:
//            - "miscellaneous"
//            - "base"
//            - "additional"
//            - "modifier"
//            - "setup_fee"
// required:
//    - "id"
//    - "name"
//    - "isRecurring"
//    - "type"
//    - "isManagedByDomain"

type ArticleTemplate struct {
	AdditionalArticles []ReadableBookableArticleOptions `json:"additionalArticles,omitempty"`
	Addons             []ArticleAddons                  `json:"addons,omitempty"`
	Attributes         []ArticleAttributes              `json:"attributes,omitempty"`
	Description        *string                          `json:"description,omitempty"`
	Id                 string                           `json:"id"`
	IsManagedByDomain  bool                             `json:"isManagedByDomain"`
	IsRecurring        bool                             `json:"isRecurring"`
	ModifierArticles   []ReadableModifierArticleOptions `json:"modifierArticles,omitempty"`
	Name               string                           `json:"name"`
	Type               ArticleTemplateType              `json:"type"`
}

func (o *ArticleTemplate) Validate() error {
	if err := func() error {
		if o.AdditionalArticles == nil {
			return nil
		}
		return func() error {
			for i := range o.AdditionalArticles {
				if err := o.AdditionalArticles[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property additionalArticles: %w", err)
	}
	if err := func() error {
		if o.Addons == nil {
			return nil
		}
		return func() error {
			for i := range o.Addons {
				if err := o.Addons[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property addons: %w", err)
	}
	if err := func() error {
		if o.Attributes == nil {
			return nil
		}
		return func() error {
			for i := range o.Attributes {
				if err := o.Attributes[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property attributes: %w", err)
	}
	if err := func() error {
		if o.ModifierArticles == nil {
			return nil
		}
		return func() error {
			for i := range o.ModifierArticles {
				if err := o.ModifierArticles[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property modifierArticles: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}
