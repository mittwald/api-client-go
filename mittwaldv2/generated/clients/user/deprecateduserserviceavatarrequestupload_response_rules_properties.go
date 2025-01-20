package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "imageDimensions":
//        type: "object"
//        properties:
//            "max":
//                type: "object"
//                properties:
//                    "height":
//                        type: "integer"
//                    "width":
//                        type: "integer"
//            "min":
//                type: "object"
//                properties:
//                    "height":
//                        type: "integer"
//                    "width":
//                        type: "integer"

//
type DeprecatedUserServiceAvatarRequestUploadResponseRulesProperties struct {
	ImageDimensions *DeprecatedUserServiceAvatarRequestUploadResponseRulesPropertiesImageDimensions `json:"imageDimensions,omitempty"`
}

func (o *DeprecatedUserServiceAvatarRequestUploadResponseRulesProperties) Validate() error {
	if err := func() error {
		if o.ImageDimensions == nil {
			return nil
		}
		return o.ImageDimensions.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property imageDimensions: %w", err)
	}
	return nil
}