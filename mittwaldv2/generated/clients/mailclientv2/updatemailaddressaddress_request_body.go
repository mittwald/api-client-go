package mailclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "address":
//        type: "string"
//        format: "idn-email"
// required:
//    - "address"
// description: UpdateMailAddressAddressRequestBody models the JSON body of a 'mail-update-mail-address-address' request

// UpdateMailAddressAddressRequestBody models the JSON body of a 'mail-update-mail-address-address' request
type UpdateMailAddressAddressRequestBody struct {
	Address string `json:"address"`
}

func (o *UpdateMailAddressAddressRequestBody) Validate() error {
	return nil
}
