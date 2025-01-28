package file

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "inline"
//    - "attachment"
// default: "inline"
// example: "inline"

type GetFileRequestHeaderContentDisposition string

const GetFileRequestHeaderContentDispositionInline GetFileRequestHeaderContentDisposition = "inline"
const GetFileRequestHeaderContentDispositionAttachment GetFileRequestHeaderContentDisposition = "attachment"

func (e GetFileRequestHeaderContentDisposition) Validate() error {
	if e == GetFileRequestHeaderContentDispositionInline || e == GetFileRequestHeaderContentDispositionAttachment {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
