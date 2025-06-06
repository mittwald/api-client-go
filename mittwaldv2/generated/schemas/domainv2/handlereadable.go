package domainv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "current": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleData"}
//    "desired": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleData"}
// required:
//    - "current"

type HandleReadable struct {
	Current HandleData  `json:"current"`
	Desired *HandleData `json:"desired,omitempty"`
}

func (o *HandleReadable) Validate() error {
	if err := o.Current.Validate(); err != nil {
		return fmt.Errorf("invalid property current: %w", err)
	}
	if err := func() error {
		if o.Desired == nil {
			return nil
		}
		return o.Desired.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property desired: %w", err)
	}
	return nil
}
