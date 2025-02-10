package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// properties:
//    "notificationThresholdInBytes":
//        type: "integer"
//        nullable: true
//        example: 10000
// description: StoragespaceReplaceServerNotificationThresholdRequestBody models the JSON body of a 'storagespace-replace-server-notification-threshold' request

// StoragespaceReplaceServerNotificationThresholdRequestBody models the JSON body of a 'storagespace-replace-server-notification-threshold' request
type StoragespaceReplaceServerNotificationThresholdRequestBody struct {
	NotificationThresholdInBytes *int64 `json:"notificationThresholdInBytes,omitempty"`
}

func (o *StoragespaceReplaceServerNotificationThresholdRequestBody) Validate() error {
	return nil
}
