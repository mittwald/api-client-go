package file

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "application/octet-stream"
//    - "text/plain;base64"
//default: "application/octet-stream"
//example: "application/octet-stream"

type GetFileWithNameRequestHeaderAccept string

const GetFileWithNameRequestHeaderAcceptApplicationOctetStream GetFileWithNameRequestHeaderAccept = "application/octet-stream"
const GetFileWithNameRequestHeaderAcceptTextPlainBase64 GetFileWithNameRequestHeaderAccept = "text/plain;base64"

func (e GetFileWithNameRequestHeaderAccept) Validate() error {
	if e == GetFileWithNameRequestHeaderAcceptApplicationOctetStream || e == GetFileWithNameRequestHeaderAcceptTextPlainBase64 {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}