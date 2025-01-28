package dnsv1

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "records":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSRVRecord"}
//        minItems: 1
//    "settings": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSettings"}
// required:
//    - "settings"
//    - "records"
// additionalProperties: false

type RecordSRVComponent struct {
	Records  []RecordSRVRecord `json:"records"`
	Settings RecordSettings    `json:"settings"`
}

func (o *RecordSRVComponent) Validate() error {
	if o.Records == nil {
		return errors.New("property records is required, but not set")
	}
	if err := func() error {
		for i := range o.Records {
			if err := o.Records[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property records: %w", err)
	}
	if err := o.Settings.Validate(); err != nil {
		return fmt.Errorf("invalid property settings: %w", err)
	}
	return nil
}
