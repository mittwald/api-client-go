package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "quotaInBytes":
//        type: "number"
//        minimum: -1
//        description: "2 GB"
//        example: 2147483648
//required:
//    - "quotaInBytes"

//
type UpdateMailAddressQuotaV2DeprecatedRequestBody struct {
	QuotaInBytes float64 `json:"quotaInBytes"`
}

func (o *UpdateMailAddressQuotaV2DeprecatedRequestBody) Validate() error {
	return nil
}
