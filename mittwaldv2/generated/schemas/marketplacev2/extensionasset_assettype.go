package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "image"
//    - "video"
// example: "image"

type ExtensionAssetAssetType string

const ExtensionAssetAssetTypeImage ExtensionAssetAssetType = "image"
const ExtensionAssetAssetTypeVideo ExtensionAssetAssetType = "video"

func (e ExtensionAssetAssetType) Validate() error {
	if e == ExtensionAssetAssetTypeImage || e == ExtensionAssetAssetTypeVideo {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
