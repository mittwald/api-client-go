package projectv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "height":
//        type: "integer"
//    "width":
//        type: "integer"

type AvatarRulesPropertiesImageDimensionsMin struct {
	Height *int64 `json:"height,omitempty"`
	Width  *int64 `json:"width,omitempty"`
}

func (o *AvatarRulesPropertiesImageDimensionsMin) Validate() error {
	return nil
}
