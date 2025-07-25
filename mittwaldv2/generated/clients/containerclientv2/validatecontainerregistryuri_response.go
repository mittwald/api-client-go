package containerclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "reason":
//        type: "string"
//    "valid":
//        type: "boolean"
// required:
//    - "valid"

type ValidateContainerRegistryUriResponse struct {
	Reason *string `json:"reason,omitempty"`
	Valid  bool    `json:"valid"`
}

func (o *ValidateContainerRegistryUriResponse) Validate() error {
	return nil
}
