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

type GetFileRequestQueryAccept string

const GetFileRequestQueryAcceptApplicationOctetStream GetFileRequestQueryAccept = "application/octet-stream"
const GetFileRequestQueryAcceptTextPlainBase64 GetFileRequestQueryAccept = "text/plain;base64"

func (e GetFileRequestQueryAccept) Validate() error {
	if e == GetFileRequestQueryAcceptApplicationOctetStream || e == GetFileRequestQueryAcceptTextPlainBase64 {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}