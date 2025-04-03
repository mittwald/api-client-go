package containerv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "credentials": {"$ref": "#/components/schemas/de.mittwald.v1.container.RegistryCredentials"}
//    "description":
//        type: "string"
//        example: "DockerHub"
//    "id":
//        type: "string"
//        format: "uuid"
//    "projectId":
//        type: "string"
//        format: "uuid"
//    "uri":
//        type: "string"
//        example: "index.docker.io"
// required:
//    - "id"
//    - "projectId"
//    - "uri"
//    - "description"

type Registry struct {
	Credentials *RegistryCredentials `json:"credentials,omitempty"`
	Description string               `json:"description"`
	Id          string               `json:"id"`
	ProjectId   string               `json:"projectId"`
	Uri         string               `json:"uri"`
}

func (o *Registry) Validate() error {
	if err := func() error {
		if o.Credentials == nil {
			return nil
		}
		return o.Credentials.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property credentials: %w", err)
	}
	return nil
}
