package project

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//        description: "Name of the Project as it is displayed in the mStudio."
//        example: "My first Project!"
// required:
//    - "description"
// description: CreateProjectRequestBody models the JSON body of a 'project-create-project' request

// CreateProjectRequestBody models the JSON body of a 'project-create-project' request
type CreateProjectRequestBody struct {
	Description string `json:"description"`
}

func (o *CreateProjectRequestBody) Validate() error {
	return nil
}
