package domainv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "adminC": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleReadable"}
//    "ownerC": {"$ref": "#/components/schemas/de.mittwald.v1.domain.HandleReadable"}
//required:
//    - "ownerC"

//
type DomainHandles struct {
	AdminC *HandleReadable `json:"adminC,omitempty"`
	OwnerC HandleReadable  `json:"ownerC"`
}

func (o *DomainHandles) Validate() error {
	if err := func() error {
		if o.AdminC == nil {
			return nil
		}
		return o.AdminC.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property adminC: %w", err)
	}
	if err := o.OwnerC.Validate(); err != nil {
		return fmt.Errorf("invalid property ownerC: %w", err)
	}
	return nil
}
