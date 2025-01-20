package cronjobv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "interpreter":
//        type: "string"
//        example: "/usr/bin/bash"
//    "parameters":
//        type: "string"
//        example: "--debug"
//    "path":
//        type: "string"
//        example: "/html/my-wordpress/script.sh"
//required:
//    - "interpreter"
//    - "path"

//
type CronjobCommand struct {
	Interpreter string  `json:"interpreter"`
	Parameters  *string `json:"parameters,omitempty"`
	Path        string  `json:"path"`
}

func (o *CronjobCommand) Validate() error {
	return nil
}
