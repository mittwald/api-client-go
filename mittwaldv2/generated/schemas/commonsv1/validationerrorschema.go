package commonsv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "message":
//        type: "string"
//        description: "The standard error message"
//        example: "should be string"
//    "path":
//        type: "string"
//        description: "The path to the part of the data that was validated. JavaScript property access notation (e.g., \".prop[1].subProp\") is used.\n"
//        example: ".address.street"
//    "type":
//        type: "string"
//        description: "ajv validation error type (e.g. \"format\", \"required\", \"type\") or own validation error types"
//        example: "format"
//    "context":
//        type: "object"
//        additionalProperties:
//            type: "string"
//        description: "The object with the additional information about the error that can be used to create custom error messages. Keys depend on the\ntype that failed validation (e.g. \"missingProperty\" for type \"required\")\n"
//        example: {"format": "email"}
//required:
//    - "message"
//    - "path"
//    - "type"

//
type ValidationErrorSchema struct {
	Message string            `json:"message"`
	Path    string            `json:"path"`
	Type    string            `json:"type"`
	Context map[string]string `json:"context,omitempty"`
}

func (o *ValidationErrorSchema) Validate() error {
	return nil
}
