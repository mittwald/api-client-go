package cronjob

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "appId":
//        type: "string"
//        format: "uuid"
//required:
//    - "appId"

type UpdateCronjobAppIDRequestBody struct {
	AppId uuid.UUID `json:"appId"`
}

func (o *UpdateCronjobAppIDRequestBody) Validate() error {
	return nil
}
