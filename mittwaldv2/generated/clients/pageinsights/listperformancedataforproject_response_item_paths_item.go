package pageinsights

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "path":
//        type: "string"
//    "performanceScore":
//        type: "integer"
//    "screenshotFileRef":
//        type: "string"
//required:
//    - "path"
//    - "performanceScore"
//    - "createdAt"

//
type ListPerformanceDataForProjectResponseItemPathsItem struct {
	CreatedAt         string  `json:"createdAt"`
	Path              string  `json:"path"`
	PerformanceScore  int64   `json:"performanceScore"`
	ScreenshotFileRef *string `json:"screenshotFileRef,omitempty"`
}

func (o *ListPerformanceDataForProjectResponseItemPathsItem) Validate() error {
	return nil
}
