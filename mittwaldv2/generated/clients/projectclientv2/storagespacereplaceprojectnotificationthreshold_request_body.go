package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// properties:
//    "notificationThresholdInBytes":
//        type: "integer"
//        format: "int64"
//        nullable: true
//        example: 10000
// description: StoragespaceReplaceProjectNotificationThresholdRequestBody models the JSON body of a 'storagespace-replace-project-notification-threshold' request

// StoragespaceReplaceProjectNotificationThresholdRequestBody models the JSON body of a 'storagespace-replace-project-notification-threshold' request
type StoragespaceReplaceProjectNotificationThresholdRequestBody struct {
	NotificationThresholdInBytes *int64 `json:"notificationThresholdInBytes,omitempty"`
}

func (o *StoragespaceReplaceProjectNotificationThresholdRequestBody) Validate() error {
	return nil
}
