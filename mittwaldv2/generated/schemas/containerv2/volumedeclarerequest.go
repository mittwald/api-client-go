package containerv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "name":
//        type: "string"
//        example: "mysql-volume"
// required:
//    - "name"

type VolumeDeclareRequest struct {
	Name string `json:"name"`
}

func (o *VolumeDeclareRequest) Validate() error {
	return nil
}
