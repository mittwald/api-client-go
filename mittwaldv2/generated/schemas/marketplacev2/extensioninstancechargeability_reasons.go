package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "isOwnExtension":
//        type: "boolean"
// required:
//    - "isOwnExtension"

type ExtensionInstanceChargeabilityReasons struct {
	IsOwnExtension bool `json:"isOwnExtension"`
}

func (o *ExtensionInstanceChargeabilityReasons) Validate() error {
	return nil
}
