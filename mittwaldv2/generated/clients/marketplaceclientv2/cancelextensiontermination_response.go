package marketplaceclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "extensionInstanceId":
//        type: "string"
//        format: "uuid"
// required:
//    - "extensionInstanceId"

type CancelExtensionTerminationResponse struct {
	ExtensionInstanceId string `json:"extensionInstanceId"`
}

func (o *CancelExtensionTerminationResponse) Validate() error {
	return nil
}
