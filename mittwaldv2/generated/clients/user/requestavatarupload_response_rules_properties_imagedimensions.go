package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "max":
//        type: "object"
//        properties:
//            "height":
//                type: "integer"
//            "width":
//                type: "integer"
//    "min":
//        type: "object"
//        properties:
//            "height":
//                type: "integer"
//            "width":
//                type: "integer"
//description: "Supported range of dimensions for the avatar image."

//Supported range of dimensions for the avatar image.
type RequestAvatarUploadResponseRulesPropertiesImageDimensions struct {
	Max *RequestAvatarUploadResponseRulesPropertiesImageDimensionsMax `json:"max,omitempty"`
	Min *RequestAvatarUploadResponseRulesPropertiesImageDimensionsMin `json:"min,omitempty"`
}

func (o *RequestAvatarUploadResponseRulesPropertiesImageDimensions) Validate() error {
	if err := func() error {
		if o.Max == nil {
			return nil
		}
		return o.Max.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property max: %w", err)
	}
	if err := func() error {
		if o.Min == nil {
			return nil
		}
		return o.Min.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property min: %w", err)
	}
	return nil
}
