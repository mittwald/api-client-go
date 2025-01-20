package storagespacev1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "isExceeding":
//        type: "boolean"
//        example: false
//    "limitInBytes":
//        type: "integer"
//        example: 100000
//    "notificationThresholdUsedAsLimit":
//        type: "boolean"
//        description: "If true, set notification threshold is used as limit for meta calculations. E.g. for projects with a parent server."
//        example: false
//    "totalExceedanceInBytes":
//        type: "integer"
//        example: 10
//    "totalExceedanceInBytesSetAt":
//        type: "string"
//        format: "date-time"
//        example: "2023-12-22T13:46:52.000Z"
//    "totalFreeInBytes":
//        type: "integer"
//        example: 99000
//    "totalFreeInPercentage":
//        type: "number"
//        example: 90
//    "totalUsageInBytes":
//        type: "integer"
//        example: 1000
//    "totalUsageInPercentage":
//        type: "number"
//        example: 10
//required:
//    - "totalUsageInBytes"

//
type StatisticsMeta struct {
	IsExceeding                      *bool      `json:"isExceeding,omitempty"`
	LimitInBytes                     *int64     `json:"limitInBytes,omitempty"`
	NotificationThresholdUsedAsLimit *bool      `json:"notificationThresholdUsedAsLimit,omitempty"`
	TotalExceedanceInBytes           *int64     `json:"totalExceedanceInBytes,omitempty"`
	TotalExceedanceInBytesSetAt      *time.Time `json:"totalExceedanceInBytesSetAt,omitempty"`
	TotalFreeInBytes                 *int64     `json:"totalFreeInBytes,omitempty"`
	TotalFreeInPercentage            *float64   `json:"totalFreeInPercentage,omitempty"`
	TotalUsageInBytes                int64      `json:"totalUsageInBytes"`
	TotalUsageInPercentage           *float64   `json:"totalUsageInPercentage,omitempty"`
}

func (o *StatisticsMeta) Validate() error {
	return nil
}
