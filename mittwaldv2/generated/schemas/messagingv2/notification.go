package messagingv2

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
//    "id":
//        type: "string"
//        format: "uuid"
//    "read":
//        type: "boolean"
//    "reference": {"$ref": "#/components/schemas/de.mittwald.v1.messaging.AggregateReference"}
//    "severity":
//        type: "string"
//        enum:
//            - "success"
//            - "info"
//            - "warning"
//            - "error"
//    "type":
//        type: "string"
// required:
//    - "id"
//    - "type"
//    - "reference"
//    - "severity"
//    - "read"
//    - "createdAt"

type Notification struct {
	CreatedAt time.Time            `json:"createdAt"`
	Id        string               `json:"id"`
	Read      bool                 `json:"read"`
	Reference AggregateReference   `json:"reference"`
	Severity  NotificationSeverity `json:"severity"`
	Type      string               `json:"type"`
}

func (o *Notification) Validate() error {
	if err := o.Reference.Validate(); err != nil {
		return fmt.Errorf("invalid property reference: %w", err)
	}
	if err := o.Severity.Validate(); err != nil {
		return fmt.Errorf("invalid property severity: %w", err)
	}
	return nil
}
