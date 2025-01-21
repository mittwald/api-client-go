package domainv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "handleFields":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleField"}
//    "handleRef":
//        type: "string"

type HandleData struct {
	HandleFields []HandleField `json:"handleFields,omitempty"`
	HandleRef    *string       `json:"handleRef,omitempty"`
}

func (o *HandleData) Validate() error {
	if err := func() error {
		if o.HandleFields == nil {
			return nil
		}
		return func() error {
			for i := range o.HandleFields {
				if err := o.HandleFields[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property handleFields: %w", err)
	}
	return nil
}
