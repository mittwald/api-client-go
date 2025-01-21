package customerv1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "createdAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "createdAt"

type CustomerActiveSuspension struct {
	CreatedAt time.Time `json:"createdAt"`
}

func (o *CustomerActiveSuspension) Validate() error {
	return nil
}
