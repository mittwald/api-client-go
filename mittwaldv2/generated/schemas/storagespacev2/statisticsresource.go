package storagespacev2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//        example: "MySQL DB for Wordpress"
//    "id":
//        type: "string"
//        example: "169cea81-2c11-46a4-8f0b-5b0b47caeb78"
//    "name":
//        type: "string"
//        example: "mysql-xyz"
//    "usageInBytes":
//        type: "integer"
//        format: "int64"
//        example: 1000
//    "usageInBytesSetAt":
//        type: "string"
//        format: "date-time"
//        example: "2023-12-22T13:46:52.000Z"
// required:
//    - "id"
//    - "name"
//    - "usageInBytes"
//    - "usageInBytesSetAt"

type StatisticsResource struct {
	Description       *string   `json:"description,omitempty"`
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	UsageInBytes      int64     `json:"usageInBytes"`
	UsageInBytesSetAt time.Time `json:"usageInBytesSetAt"`
}

func (o *StatisticsResource) Validate() error {
	return nil
}
