package pageinsightsclientv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "fileRef":
//        type: "string"
// required:
//    - "fileRef"
//    - "createdAt"

type GetPerformanceDataResponseScreenshot struct {
	CreatedAt time.Time `json:"createdAt"`
	FileRef   string    `json:"fileRef"`
}

func (o *GetPerformanceDataResponseScreenshot) Validate() error {
	return nil
}
