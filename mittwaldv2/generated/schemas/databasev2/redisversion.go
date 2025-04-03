package databasev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "disabled":
//        type: "boolean"
//    "id":
//        type: "string"
//    "name":
//        type: "string"
//    "number":
//        type: "string"
// required:
//    - "id"
//    - "name"
//    - "number"
//    - "disabled"

type RedisVersion struct {
	Disabled bool   `json:"disabled"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Number   string `json:"number"`
}

func (o *RedisVersion) Validate() error {
	return nil
}
