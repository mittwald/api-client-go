package dnsv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "managed":
//        type: "boolean"
// required:
//    - "managed"

type RecordMXManaged struct {
	Managed bool `json:"managed"`
}

func (o *RecordMXManaged) Validate() error {
	return nil
}
