package varnishv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "maxVersions":
//        type: "number"
//    "retentionTime":
//        type: "number"

type ConfigExpiration struct {
	MaxVersions   *float64 `json:"maxVersions,omitempty"`
	RetentionTime *float64 `json:"retentionTime,omitempty"`
}

func (o *ConfigExpiration) Validate() error {
	return nil
}
