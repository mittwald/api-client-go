package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "description":
//        type: "string"
//        example: "My first Server!"
// required:
//    - "description"
// description: UpdateServerDescriptionRequestBody models the JSON body of a 'project-update-server-description' request

// UpdateServerDescriptionRequestBody models the JSON body of a 'project-update-server-description' request
type UpdateServerDescriptionRequestBody struct {
	Description string `json:"description"`
}

func (o *UpdateServerDescriptionRequestBody) Validate() error {
	return nil
}
