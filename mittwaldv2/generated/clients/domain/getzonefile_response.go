package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "file":
//        type: "string"

type GetZoneFileResponse struct {
	File *string `json:"file,omitempty"`
}

func (o *GetZoneFileResponse) Validate() error {
	return nil
}
