package invoicev1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "end":
//        type: "string"
//        format: "date-time"
//    "start":
//        type: "string"
//        format: "date-time"
//required:
//    - "start"
//    - "end"

type DatePeriod struct {
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

func (o *DatePeriod) Validate() error {
	return nil
}
