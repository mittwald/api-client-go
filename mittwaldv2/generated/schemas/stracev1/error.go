package stracev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "errorMessage":
//        type: "string"
//        example: "wrong PHP version"
//required:
//    - "errorMessage"

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

func (o *Error) Validate() error {
	return nil
}
