package marketplace

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "extensionInstanceSecret":
//        type: "string"
//        description: "The latest secret an external application received via lifecycle webhooks.\nNamely ExtensionAddedToContext and ExtensionInstanceSecretRotated.\n"
//required:
//    - "extensionInstanceSecret"

type AuthenticateInstanceRequestBody struct {
	ExtensionInstanceSecret string `json:"extensionInstanceSecret"`
}

func (o *AuthenticateInstanceRequestBody) Validate() error {
	return nil
}
