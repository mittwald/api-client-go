package marketplacev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "name":
//        type: "string"
//    "url":
//        type: "string"
//required:
//    - "name"
//    - "url"

type ExternalComponent struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (o *ExternalComponent) Validate() error {
	return nil
}
