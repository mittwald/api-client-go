package domainclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "isPremium":
//        type: "boolean"
//    "registrable":
//        type: "boolean"
// required:
//    - "registrable"
//    - "isPremium"

type CheckDomainRegistrabilityResponse struct {
	IsPremium   bool `json:"isPremium"`
	Registrable bool `json:"registrable"`
}

func (o *CheckDomainRegistrabilityResponse) Validate() error {
	return nil
}
