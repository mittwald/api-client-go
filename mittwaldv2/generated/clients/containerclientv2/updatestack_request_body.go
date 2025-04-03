package containerclientv2

import "github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/containerv2"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "services":
//        type: "object"
//        additionalProperties: {"$ref": "#/components/schemas/de.mittwald.v1.container.ServiceRequest"}
//    "volumes":
//        type: "object"
//        additionalProperties: {"$ref": "#/components/schemas/de.mittwald.v1.container.VolumeRequest"}
//        description: "Volumes belonging to a Stack. Removing results in a detach, delete must be explicit."
// description: UpdateStackRequestBody models the JSON body of a 'container-update-stack' request

// UpdateStackRequestBody models the JSON body of a 'container-update-stack' request
type UpdateStackRequestBody struct {
	Services map[string]containerv2.ServiceRequest `json:"services,omitempty"`
	Volumes  map[string]containerv2.VolumeRequest  `json:"volumes,omitempty"`
}

func (o *UpdateStackRequestBody) Validate() error {
	return nil
}
