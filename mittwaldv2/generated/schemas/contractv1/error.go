package contractv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "message":
//        type: "string"
//        example: "Something went wrong"
//    "type":
//        type: "string"
//        example: "InternalServerError"
//required:
//    - "type"
//    - "message"

//
type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (o *Error) Validate() error {
	return nil
}
