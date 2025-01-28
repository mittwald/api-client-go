package taskv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "aggregate":
//        type: "string"
//    "domain":
//        type: "string"
//    "id":
//        type: "string"
// required:
//    - "id"
//    - "aggregate"
//    - "domain"

type AggregateReference struct {
	Aggregate string `json:"aggregate"`
	Domain    string `json:"domain"`
	Id        string `json:"id"`
}

func (o *AggregateReference) Validate() error {
	return nil
}
