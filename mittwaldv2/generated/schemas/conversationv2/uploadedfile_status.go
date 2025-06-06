package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "uploaded"

type UploadedFileStatus string

const UploadedFileStatusUploaded UploadedFileStatus = "uploaded"

func (e UploadedFileStatus) Validate() error {
	if e == UploadedFileStatusUploaded {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
