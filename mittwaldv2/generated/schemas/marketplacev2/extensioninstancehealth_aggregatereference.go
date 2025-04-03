package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "aggregate":
//        type: "string"
//    "domain":
//        type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
// required:
//    - "id"
//    - "domain"
//    - "aggregate"

type ExtensionInstanceHealthAggregateReference struct {
	Aggregate string `json:"aggregate"`
	Domain    string `json:"domain"`
	Id        string `json:"id"`
}

func (o *ExtensionInstanceHealthAggregateReference) Validate() error {
	return nil
}
