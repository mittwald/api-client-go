package marketplaceclientv2

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "extensions":
//        type: "array"
//        items:
//            type: "string"
//            example: ".png"
//    "fileTypes":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "extensions":
//                    type: "array"
//                    items:
//                        type: "string"
//                        example: ".png"
//                "mimeType":
//                    type: "string"
//                    example: "image/png"
//            required:
//                - "mimeType"
//                - "extensions"
//    "maxSizeInBytes":
//        type: "integer"
//    "mimeTypes":
//        type: "array"
//        items:
//            type: "string"
//            example: "image/png"
//    "properties":
//        type: "object"
//        properties:
//            "imageDimensions":
//                type: "object"
//                properties:
//                    "max":
//                        type: "object"
//                        properties:
//                            "height":
//                                type: "integer"
//                            "width":
//                                type: "integer"
//                    "min":
//                        type: "object"
//                        properties:
//                            "height":
//                                type: "integer"
//                            "width":
//                                type: "integer"
//                required:
//                    - "min"
//                    - "max"
// required:
//    - "mimeTypes"
//    - "extensions"
//    - "fileTypes"
//    - "maxSizeInBytes"
// description: "Constraints for the logo image upload."

// Constraints for the logo image upload.
type RequestLogoUploadResponseRules struct {
	Extensions     []string                                      `json:"extensions"`
	FileTypes      []RequestLogoUploadResponseRulesFileTypesItem `json:"fileTypes"`
	MaxSizeInBytes int64                                         `json:"maxSizeInBytes"`
	MimeTypes      []string                                      `json:"mimeTypes"`
	Properties     *RequestLogoUploadResponseRulesProperties     `json:"properties,omitempty"`
}

func (o *RequestLogoUploadResponseRules) Validate() error {
	if o.Extensions == nil {
		return errors.New("property extensions is required, but not set")
	}
	if o.FileTypes == nil {
		return errors.New("property fileTypes is required, but not set")
	}
	if err := func() error {
		for i := range o.FileTypes {
			if err := o.FileTypes[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property fileTypes: %w", err)
	}
	if o.MimeTypes == nil {
		return errors.New("property mimeTypes is required, but not set")
	}
	if err := func() error {
		if o.Properties == nil {
			return nil
		}
		return o.Properties.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property properties: %w", err)
	}
	return nil
}
