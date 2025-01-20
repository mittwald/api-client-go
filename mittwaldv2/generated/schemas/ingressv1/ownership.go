package ingressv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "txtRecord":
//        type: "string"
//    "verified":
//        type: "boolean"
//        description: "Whether the domain ownership is verified or not."
//required:
//    - "verified"

//
type Ownership struct {
	TxtRecord *string `json:"txtRecord,omitempty"`
	Verified  bool    `json:"verified"`
}

func (o *Ownership) Validate() error {
	return nil
}