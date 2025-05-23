package dnsv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "flags":
//        type: "integer"
//        maximum: 255
//        minimum: 0
//    "tag":
//        type: "string"
//        enum:
//            - "issue"
//            - "issuewild"
//            - "iodef"
//    "value":
//        type: "string"
//        format: "idn-dnsname"
// required:
//    - "flags"
//    - "tag"
//    - "value"
// additionalProperties: false

type RecordCAARecord struct {
	Flags int64              `json:"flags"`
	Tag   RecordCAARecordTag `json:"tag"`
	Value string             `json:"value"`
}

func (o *RecordCAARecord) Validate() error {
	if err := o.Tag.Validate(); err != nil {
		return fmt.Errorf("invalid property tag: %w", err)
	}
	return nil
}
