package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "address":
//        type: "string"
//        format: "idn-email"
//required:
//    - "address"

type MailaddressUpdateAddressDeprecatedRequestBody struct {
	Address string `json:"address"`
}

func (o *MailaddressUpdateAddressDeprecatedRequestBody) Validate() error {
	return nil
}
