package containerclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "registryUri":
//        type: "string"
// required:
//    - "registryUri"
// description: "The Registry URI to validate."

// The Registry URI to validate.
type ValidateContainerRegistryUriRequestBody struct {
	RegistryUri string `json:"registryUri"`
}

func (o *ValidateContainerRegistryUriRequestBody) Validate() error {
	return nil
}
