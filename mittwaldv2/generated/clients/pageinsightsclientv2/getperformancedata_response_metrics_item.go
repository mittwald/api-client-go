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
//    "name":
//        type: "string"
//    "score":
//        type: "number"
//        format: "double"
//        nullable: true
//    "value":
//        type: "number"
//        format: "double"
// required:
//    - "name"
//    - "value"
//    - "createdAt"

type GetPerformanceDataResponseMetricsItem struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Score     *float64  `json:"score,omitempty"`
	Value     float64   `json:"value"`
}

func (o *GetPerformanceDataResponseMetricsItem) Validate() error {
	return nil
}
