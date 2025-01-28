package user

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "maxSizeInBytes":
//        type: "integer"
//        description: "Maximum size in Bytes of the avatar image."
//        example: 4096
//    "maxSizeInKB":
//        type: "integer"
//        description: "Deprecated. Maximum size in kilobytes of the avatar image."
//        example: 3000
//        deprecated: true
//    "mimeTypes":
//        type: "array"
//        items:
//            type: "string"
//            example: "image/png"
//        description: "List of supported mime types."
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
//                description: "Supported range of dimensions for the avatar image."
// required:
//    - "mimeTypes"
//    - "maxSizeInKB"
//    - "maxSizeInBytes"
// description: "Contstraints for the avatar image upload."

// Contstraints for the avatar image upload.
type RequestAvatarUploadResponseRules struct {
	MaxSizeInBytes int64                                       `json:"maxSizeInBytes"`
	MaxSizeInKB    int64                                       `json:"maxSizeInKB"`
	MimeTypes      []string                                    `json:"mimeTypes"`
	Properties     *RequestAvatarUploadResponseRulesProperties `json:"properties,omitempty"`
}

func (o *RequestAvatarUploadResponseRules) Validate() error {
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
