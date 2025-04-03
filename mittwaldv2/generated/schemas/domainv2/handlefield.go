package domainv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "name":
//        type: "string"
//    "value":
//        type: "string"
//        pattern: "^[0-9a-zA-ZÀ-ÖØ-öø-ÿ /&()+,.@_-]+$"
// required:
//    - "name"
//    - "value"

type HandleField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *HandleField) Validate() error {
	return nil
}
