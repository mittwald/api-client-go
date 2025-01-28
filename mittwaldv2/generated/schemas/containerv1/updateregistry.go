package containerv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "credentials":
//        allOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.container.SetRegistryCredentials"}
//        nullable: true
//    "description":
//        type: "string"
//        example: "DockerHub"
//    "uri":
//        type: "string"
//        example: "index.docker.io"

type UpdateRegistry struct {
	Credentials *any    `json:"credentials,omitempty"`
	Description *string `json:"description,omitempty"`
	Uri         *string `json:"uri,omitempty"`
}

func (o *UpdateRegistry) Validate() error {
	return nil
}
