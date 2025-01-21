package marketplacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "de": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.DescriptionFormats"}
//    "en": {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.DescriptionFormats"}
//required:
//    - "de"
//description: "Supported languages. Format ISO-639-1."

// Supported languages. Format ISO-639-1.
type DetailedDescriptions struct {
	De DescriptionFormats  `json:"de"`
	En *DescriptionFormats `json:"en,omitempty"`
}

func (o *DetailedDescriptions) Validate() error {
	if err := o.De.Validate(); err != nil {
		return fmt.Errorf("invalid property de: %w", err)
	}
	if err := func() error {
		if o.En == nil {
			return nil
		}
		return o.En.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property en: %w", err)
	}
	return nil
}
