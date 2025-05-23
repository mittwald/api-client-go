package dnsv2

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "records":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordMXRecord"}
//        maxItems: 10
//        minItems: 1
//        uniqueItems: true
//    "settings": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSettings"}
// required:
//    - "settings"
//    - "records"

type RecordMXCustom struct {
	Records  []RecordMXRecord `json:"records"`
	Settings RecordSettings   `json:"settings"`
}

func (o *RecordMXCustom) Validate() error {
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
