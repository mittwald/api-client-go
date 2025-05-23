package marketplaceclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "failure":
//        type: "boolean"
//    "statusCode":
//        type: "string"
//        description: "The status code returned by the external application."
// required:
//    - "failure"

type RotateSecretForExtensionInstanceResponseWebhookResult struct {
	Failure    bool    `json:"failure"`
	StatusCode *string `json:"statusCode,omitempty"`
}

func (o *RotateSecretForExtensionInstanceResponseWebhookResult) Validate() error {
	return nil
}
