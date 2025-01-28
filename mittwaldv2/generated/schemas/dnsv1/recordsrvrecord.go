package dnsv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "fqdn":
//        type: "string"
//        maxLength: 50
//        minLength: 1
//        format: "idn-dnsname"
//    "port":
//        type: "integer"
//        maximum: 65535
//        minimum: 0
//    "priority":
//        type: "integer"
//        maximum: 65535
//        minimum: 0
//    "weight":
//        type: "integer"
//        maximum: 65535
//        minimum: 0
// required:
//    - "port"
//    - "fqdn"
// additionalProperties: false

type RecordSRVRecord struct {
	Fqdn     string `json:"fqdn"`
	Port     int64  `json:"port"`
	Priority *int64 `json:"priority,omitempty"`
	Weight   *int64 `json:"weight,omitempty"`
}

func (o *RecordSRVRecord) Validate() error {
	return nil
}
