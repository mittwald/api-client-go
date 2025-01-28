package screenshotv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "dataType":
//        type: "string"
//        enum:
//            - "jpeg"
//            - "png"
//            - "webp"
//    "delay":
//        type: "number"
//    "height":
//        type: "number"
//    "quality":
//        type: "number"
//    "width":
//        type: "number"
// required:
//    - "width"
//    - "height"
//    - "quality"
//    - "delay"
//    - "dataType"

type ScreenshotSettings struct {
	DataType ScreenshotSettingsDataType `json:"dataType"`
	Delay    float64                    `json:"delay"`
	Height   float64                    `json:"height"`
	Quality  float64                    `json:"quality"`
	Width    float64                    `json:"width"`
}

func (o *ScreenshotSettings) Validate() error {
	if err := o.DataType.Validate(); err != nil {
		return fmt.Errorf("invalid property dataType: %w", err)
	}
	return nil
}
