package marketplace

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "allowWebhookFailure":
//        type: "boolean"
// description: RotateSecretForExtensionInstanceRequestBody models the JSON body of a 'contributor-rotate-secret-for-extension-instance' request

// RotateSecretForExtensionInstanceRequestBody models the JSON body of a 'contributor-rotate-secret-for-extension-instance' request
type RotateSecretForExtensionInstanceRequestBody struct {
	AllowWebhookFailure *bool `json:"allowWebhookFailure,omitempty"`
}

func (o *RotateSecretForExtensionInstanceRequestBody) Validate() error {
	return nil
}
