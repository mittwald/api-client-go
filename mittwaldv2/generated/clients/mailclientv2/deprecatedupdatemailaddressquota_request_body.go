package mailclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "quotaInBytes":
//        type: "number"
//        minimum: -1
//        description: "2 GB"
//        example: 2147483648
// required:
//    - "quotaInBytes"
// description: DeprecatedUpdateMailAddressQuotaRequestBody models the JSON body of a 'deprecated-mail-update-mail-address-quota' request

// DeprecatedUpdateMailAddressQuotaRequestBody models the JSON body of a 'deprecated-mail-update-mail-address-quota' request
type DeprecatedUpdateMailAddressQuotaRequestBody struct {
	QuotaInBytes float64 `json:"quotaInBytes"`
}

func (o *DeprecatedUpdateMailAddressQuotaRequestBody) Validate() error {
	return nil
}
