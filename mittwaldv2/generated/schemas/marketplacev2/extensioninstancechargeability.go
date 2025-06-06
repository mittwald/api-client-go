package marketplacev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "isChargeable":
//        type: "boolean"
//        default: false
//    "reasons":
//        type: "object"
//        properties:
//            "isOwnExtension":
//                type: "boolean"
//        required:
//            - "isOwnExtension"
// required:
//    - "isChargeable"
//    - "reasons"

type ExtensionInstanceChargeability struct {
	IsChargeable bool                                  `json:"isChargeable"`
	Reasons      ExtensionInstanceChargeabilityReasons `json:"reasons"`
}

func (o *ExtensionInstanceChargeability) Validate() error {
	if err := o.Reasons.Validate(); err != nil {
		return fmt.Errorf("invalid property reasons: %w", err)
	}
	return nil
}
