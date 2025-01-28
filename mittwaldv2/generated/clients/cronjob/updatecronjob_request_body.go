package cronjob

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "active":
//        type: "boolean"
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
// description: UpdateCronjobRequestBody models the JSON body of a 'cronjob-update-cronjob' request

// UpdateCronjobRequestBody models the JSON body of a 'cronjob-update-cronjob' request
type UpdateCronjobRequestBody struct {
	Active      *bool                                `json:"active,omitempty"`
	Description *string                              `json:"description,omitempty"`
	Destination *UpdateCronjobRequestBodyDestination `json:"destination,omitempty"`
	Email       *string                              `json:"email,omitempty"`
	Interval    *string                              `json:"interval,omitempty"`
	Timeout     *float64                             `json:"timeout,omitempty"`
}

func (o *UpdateCronjobRequestBody) Validate() error {
	if err := func() error {
		if o.Destination == nil {
			return nil
		}
		return o.Destination.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property destination: %w", err)
	}
	return nil
}
