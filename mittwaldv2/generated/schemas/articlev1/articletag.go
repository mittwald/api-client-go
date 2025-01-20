package articlev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "description":
//        type: "string"
//        minLength: 1
//        example: "CPU optimized"
//    "hexColor":
//        type: "string"
//        minLength: 1
//        example: "#ff0000"
//    "id":
//        type: "string"
//        minLength: 1
//        example: "58e39bbe-fcd1-4fde-887c-f49ff85b9df5"
//    "name":
//        type: "string"
//        minLength: 1
//        example: "cpu-optimized"
//required:
//    - "id"

//
type ArticleTag struct {
	Description *string `json:"description,omitempty"`
	HexColor    *string `json:"hexColor,omitempty"`
	Id          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
}

func (o *ArticleTag) Validate() error {
	return nil
}