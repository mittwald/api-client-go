package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "available":
//        type: "boolean"
//required:
//    - "available"

//
type CheckDomainRegistrabilityV2DeprecatedResponse struct {
	Available bool `json:"available"`
}

func (o *CheckDomainRegistrabilityV2DeprecatedResponse) Validate() error {
	return nil
}