package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "reference":
//        type: "string"

type GetLatestScreenshotResponse struct {
	Reference *string `json:"reference,omitempty"`
}

func (o *GetLatestScreenshotResponse) Validate() error {
	return nil
}
