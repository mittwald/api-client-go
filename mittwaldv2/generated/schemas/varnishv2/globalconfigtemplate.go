package varnishv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "data":
//        type: "string"
//        format: "byte"
//    "name":
//        type: "string"
//    "updatedAt":
//        type: "string"
//        format: "date-time"
// required:
//    - "name"
//    - "data"
//    - "updatedAt"

type GlobalConfigTemplate struct {
	Data      string    `json:"data"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (o *GlobalConfigTemplate) Validate() error {
	return nil
}
