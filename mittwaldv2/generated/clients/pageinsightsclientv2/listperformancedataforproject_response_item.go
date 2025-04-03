package pageinsightsclientv2

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "domain":
//        type: "string"
//    "paths":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "createdAt":
//                    type: "string"
//                    format: "date-time"
//                "path":
//                    type: "string"
//                "performanceScore":
//                    type: "integer"
//                "screenshotFileRef":
//                    type: "string"
//            required:
//                - "path"
//                - "performanceScore"
//                - "createdAt"
// required:
//    - "domain"
//    - "paths"

type ListPerformanceDataForProjectResponseItem struct {
	Domain string                                               `json:"domain"`
	Paths  []ListPerformanceDataForProjectResponseItemPathsItem `json:"paths"`
}

func (o *ListPerformanceDataForProjectResponseItem) Validate() error {
	if o.Paths == nil {
		return errors.New("property paths is required, but not set")
	}
	if err := func() error {
		for i := range o.Paths {
			if err := o.Paths[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property paths: %w", err)
	}
	return nil
}
