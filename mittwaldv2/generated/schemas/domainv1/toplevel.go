package domainv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "irtp":
//        type: "boolean"
//    "rgpDays":
//        type: "integer"
//        minimum: 0
//    "tld":
//        type: "string"
//    "transferAuthCodeRequired":
//        type: "boolean"
//        deprecated: true
//    "transferAuthentication": {"$ref": "#/components/schemas/de.mittwald.v1.domain.TransferAuthentication"}
// required:
//    - "tld"
//    - "rgpDays"
//    - "irtp"
//    - "transferAuthCodeRequired"
//    - "transferAuthentication"

type TopLevel struct {
	Irtp                     bool                   `json:"irtp"`
	RgpDays                  int64                  `json:"rgpDays"`
	Tld                      string                 `json:"tld"`
	TransferAuthCodeRequired bool                   `json:"transferAuthCodeRequired"`
	TransferAuthentication   TransferAuthentication `json:"transferAuthentication"`
}

func (o *TopLevel) Validate() error {
	if err := o.TransferAuthentication.Validate(); err != nil {
		return fmt.Errorf("invalid property transferAuthentication: %w", err)
	}
	return nil
}
