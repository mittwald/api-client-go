package notificationclientv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/messagingv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.messaging.NotificationStatus"}
// required:
//    - "status"
// description: ReadNotificationRequestBody models the JSON body of a 'notifications-read-notification' request

// ReadNotificationRequestBody models the JSON body of a 'notifications-read-notification' request
type ReadNotificationRequestBody struct {
	Status messagingv2.NotificationStatus `json:"status"`
}

func (o *ReadNotificationRequestBody) Validate() error {
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	return nil
}
