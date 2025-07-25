package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "createWebuiContainer":
//        type: "boolean"
//    "name":
//        type: "string"
//        minLength: 5
// description: UpdateLlmLicenceExperimentalRequestBody models the JSON body of a 'project-update-llm-licence-experimental' request

// UpdateLlmLicenceExperimentalRequestBody models the JSON body of a 'project-update-llm-licence-experimental' request
type UpdateLlmLicenceExperimentalRequestBody struct {
	CreateWebuiContainer *bool   `json:"createWebuiContainer,omitempty"`
	Name                 *string `json:"name,omitempty"`
}

func (o *UpdateLlmLicenceExperimentalRequestBody) Validate() error {
	return nil
}
