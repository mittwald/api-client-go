package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// properties:
//    "description":
//        type: "string"
// required:
//    - "description"
// description: DeprecatedUpdateDescriptionRequestBody models the JSON body of a 'deprecated-mail-deliverybox-update-description' request

// DeprecatedUpdateDescriptionRequestBody models the JSON body of a 'deprecated-mail-deliverybox-update-description' request
type DeprecatedUpdateDescriptionRequestBody struct {
	Description string `json:"description"`
}

func (o *DeprecatedUpdateDescriptionRequestBody) Validate() error {
	return nil
}
