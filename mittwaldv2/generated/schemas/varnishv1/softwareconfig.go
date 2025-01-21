package varnishv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "configExpiration": {"$ref": "#/components/schemas/de.mittwald.v1.varnish.ConfigExpiration"}
//    "latestConfigRevision":
//        type: "number"
//    "revisions":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.varnish.ConfigRevision"}

type SoftwareConfig struct {
	ConfigExpiration     *ConfigExpiration `json:"configExpiration,omitempty"`
	LatestConfigRevision *float64          `json:"latestConfigRevision,omitempty"`
	Revisions            []ConfigRevision  `json:"revisions,omitempty"`
}

func (o *SoftwareConfig) Validate() error {
	if err := func() error {
		if o.ConfigExpiration == nil {
			return nil
		}
		return o.ConfigExpiration.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property configExpiration: %w", err)
	}
	if err := func() error {
		if o.Revisions == nil {
			return nil
		}
		return func() error {
			for i := range o.Revisions {
				if err := o.Revisions[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property revisions: %w", err)
	}
	return nil
}
