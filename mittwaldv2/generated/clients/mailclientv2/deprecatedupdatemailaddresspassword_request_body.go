package mailclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "password":
//        type: "string"
// required:
//    - "password"
// description: DeprecatedUpdateMailAddressPasswordRequestBody models the JSON body of a 'deprecated-mail-update-mail-address-password' request

// DeprecatedUpdateMailAddressPasswordRequestBody models the JSON body of a 'deprecated-mail-update-mail-address-password' request
type DeprecatedUpdateMailAddressPasswordRequestBody struct {
	Password string `json:"password"`
}

func (o *DeprecatedUpdateMailAddressPasswordRequestBody) Validate() error {
	return nil
}
