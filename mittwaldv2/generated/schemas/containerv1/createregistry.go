package containerv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "credentials": {"$ref": "#/components/schemas/de.mittwald.v1.container.SetRegistryCredentials"}
//    "description":
//        type: "string"
//        example: "DockerHub"
//    "uri":
//        type: "string"
//        example: "index.docker.io"
// required:
//    - "uri"
//    - "description"

type CreateRegistry struct {
	Credentials *SetRegistryCredentials `json:"credentials,omitempty"`
	Description string                  `json:"description"`
	Uri         string                  `json:"uri"`
}

func (o *CreateRegistry) Validate() error {
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
