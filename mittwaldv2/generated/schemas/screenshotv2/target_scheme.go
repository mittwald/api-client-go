package screenshotv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "http"
//    - "https"
// example: "https"

type TargetScheme string

const TargetSchemeHttp TargetScheme = "http"
const TargetSchemeHttps TargetScheme = "https"

func (e TargetScheme) Validate() error {
	if e == TargetSchemeHttp || e == TargetSchemeHttps {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
