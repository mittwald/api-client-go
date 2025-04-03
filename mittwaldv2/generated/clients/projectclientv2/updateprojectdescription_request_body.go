package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "description":
//        type: "string"
//        example: "My new description!"
// required:
//    - "description"
// description: UpdateProjectDescriptionRequestBody models the JSON body of a 'project-update-project-description' request

// UpdateProjectDescriptionRequestBody models the JSON body of a 'project-update-project-description' request
type UpdateProjectDescriptionRequestBody struct {
	Description string `json:"description"`
}

func (o *UpdateProjectDescriptionRequestBody) Validate() error {
	return nil
}
