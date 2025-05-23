package marketplaceclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "published":
//        type: "boolean"
// required:
//    - "id"
//    - "published"

type SetExtensionPublishedStateResponse struct {
	Id        string `json:"id"`
	Published bool   `json:"published"`
}

func (o *SetExtensionPublishedStateResponse) Validate() error {
	return nil
}
