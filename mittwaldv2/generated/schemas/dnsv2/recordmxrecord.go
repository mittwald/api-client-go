package dnsv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "fqdn":
//        type: "string"
//        maxLength: 50
//        minLength: 1
//        format: "idn-dnsname"
//    "priority":
//        type: "integer"
//        maximum: 100
//        minimum: 0
// required:
//    - "priority"
//    - "fqdn"

type RecordMXRecord struct {
	Fqdn     string `json:"fqdn"`
	Priority int64  `json:"priority"`
}

func (o *RecordMXRecord) Validate() error {
	return nil
}
