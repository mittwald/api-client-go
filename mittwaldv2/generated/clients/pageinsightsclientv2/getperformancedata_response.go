package pageinsightsclientv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "domain":
//        type: "string"
//    "metrics":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "createdAt":
//                    type: "string"
//                    format: "date-time"
//                "name":
//                    type: "string"
//                "score":
//                    type: "number"
//                    format: "double"
//                    nullable: true
//                "value":
//                    type: "number"
//                    format: "double"
//            required:
//                - "name"
//                - "value"
//                - "createdAt"
//    "moreDataAvailable":
//        type: "array"
//        items:
//            type: "string"
//            format: "date"
//    "path":
//        type: "string"
//    "performanceScore":
//        type: "number"
//        format: "double"
//    "screenshot":
//        type: "object"
//        properties:
//            "createdAt":
//                type: "string"
//                format: "date-time"
//            "fileRef":
//                type: "string"
//        required:
//            - "fileRef"
//            - "createdAt"
// required:
//    - "domain"
//    - "path"
//    - "performanceScore"

type GetPerformanceDataResponse struct {
	CreatedAt         *time.Time                              `json:"createdAt,omitempty"`
	Domain            string                                  `json:"domain"`
	Metrics           []GetPerformanceDataResponseMetricsItem `json:"metrics,omitempty"`
	MoreDataAvailable []string                                `json:"moreDataAvailable,omitempty"`
	Path              string                                  `json:"path"`
	PerformanceScore  float64                                 `json:"performanceScore"`
	Screenshot        *GetPerformanceDataResponseScreenshot   `json:"screenshot,omitempty"`
}

func (o *GetPerformanceDataResponse) Validate() error {
	if err := func() error {
		if o.Metrics == nil {
			return nil
		}
		return func() error {
			for i := range o.Metrics {
				if err := o.Metrics[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property metrics: %w", err)
	}
	if err := func() error {
		if o.MoreDataAvailable == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property moreDataAvailable: %w", err)
	}
	if err := func() error {
		if o.Screenshot == nil {
			return nil
		}
		return o.Screenshot.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property screenshot: %w", err)
	}
	return nil
}
