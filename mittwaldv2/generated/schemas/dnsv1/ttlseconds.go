package dnsv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "seconds":
//        type: "integer"
//        maximum: 86400
//        minimum: 300
//required:
//    - "seconds"

//
type TtlSeconds struct {
	Seconds int64 `json:"seconds"`
}

func (o *TtlSeconds) Validate() error {
	return nil
}