package cronjobv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "active":
//        type: "boolean"
//    "appId":
//        type: "string"
//        format: "uuid"
//    "description":
//        type: "string"
//        example: "i am a cronjob"
//    "destination":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.cronjob.CronjobUrl"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.cronjob.CronjobCommand"}
//    "email":
//        type: "string"
//        format: "email"
//    "interval":
//        type: "string"
//        example: "*/5 * * * *"
//    "timeout":
//        type: "number"
//        maximum: 86400
//        minimum: 1
// required:
//    - "appId"
//    - "description"
//    - "destination"
//    - "interval"
//    - "active"
//    - "timeout"

type CronjobRequest struct {
	Active      bool                      `json:"active"`
	AppId       string                    `json:"appId"`
	Description string                    `json:"description"`
	Destination CronjobRequestDestination `json:"destination"`
	Email       *string                   `json:"email,omitempty"`
	Interval    string                    `json:"interval"`
	Timeout     float64                   `json:"timeout"`
}

func (o *CronjobRequest) Validate() error {
	if err := o.Destination.Validate(); err != nil {
		return fmt.Errorf("invalid property destination: %w", err)
	}
	return nil
}
