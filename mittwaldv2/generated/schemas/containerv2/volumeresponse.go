package containerv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "name":
//        type: "string"
//    "orphaned":
//        type: "boolean"
//        description: "Whether the Volume is attached to a Stack."
//    "stackId":
//        type: "string"
//        format: "uuid"
//    "storageUsageInBytes":
//        type: "integer"
//    "storageUsageInBytesSetAt":
//        type: "string"
//        format: "date-time"
// required:
//    - "id"
//    - "stackId"
//    - "name"
//    - "orphaned"
//    - "storageUsageInBytes"
//    - "storageUsageInBytesSetAt"

type VolumeResponse struct {
	Id                       string    `json:"id"`
	Name                     string    `json:"name"`
	Orphaned                 bool      `json:"orphaned"`
	StackId                  string    `json:"stackId"`
	StorageUsageInBytes      int64     `json:"storageUsageInBytes"`
	StorageUsageInBytesSetAt time.Time `json:"storageUsageInBytesSetAt"`
}

func (o *VolumeResponse) Validate() error {
	return nil
}
