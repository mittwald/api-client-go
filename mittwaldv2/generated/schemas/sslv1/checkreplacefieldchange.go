package sslv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "newValue":
//        type: "string"
//    "oldValue":
//        type: "string"
// required:
//    - "oldValue"
//    - "newValue"

type CheckReplaceFieldChange struct {
	NewValue string `json:"newValue"`
	OldValue string `json:"oldValue"`
}

func (o *CheckReplaceFieldChange) Validate() error {
	return nil
}
