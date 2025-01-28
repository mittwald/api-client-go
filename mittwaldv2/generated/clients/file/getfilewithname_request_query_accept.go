package file

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "application/octet-stream"
//    - "text/plain;base64"
// default: "application/octet-stream"
// example: "application/octet-stream"

type GetFileWithNameRequestQueryAccept string

const GetFileWithNameRequestQueryAcceptApplicationOctetStream GetFileWithNameRequestQueryAccept = "application/octet-stream"
const GetFileWithNameRequestQueryAcceptTextPlainBase64 GetFileWithNameRequestQueryAccept = "text/plain;base64"

func (e GetFileWithNameRequestQueryAccept) Validate() error {
	if e == GetFileWithNameRequestQueryAcceptApplicationOctetStream || e == GetFileWithNameRequestQueryAcceptTextPlainBase64 {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
