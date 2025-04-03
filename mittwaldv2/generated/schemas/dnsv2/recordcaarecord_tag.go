package dnsv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "issue"
//    - "issuewild"
//    - "iodef"

type RecordCAARecordTag string

const RecordCAARecordTagIssue RecordCAARecordTag = "issue"
const RecordCAARecordTagIssuewild RecordCAARecordTag = "issuewild"
const RecordCAARecordTagIodef RecordCAARecordTag = "iodef"

func (e RecordCAARecordTag) Validate() error {
	if e == RecordCAARecordTagIssue || e == RecordCAARecordTagIssuewild || e == RecordCAARecordTagIodef {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
