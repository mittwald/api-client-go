package databasev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "configuration": {"$ref": "#/components/schemas/de.mittwald.v1.database.RedisDatabaseConfiguration"}
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "description":
//        type: "string"
//    "finalizers":
//        type: "array"
//        items:
//            type: "string"
//            minLength: 1
//        uniqueItems: true
//    "hostname":
//        type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "name":
//        type: "string"
//    "port":
//        type: "integer"
//    "projectId":
//        type: "string"
//        format: "uuid"
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.database.DatabaseStatus"}
//    "statusSetAt":
//        type: "string"
//        format: "date-time"
//    "storageUsageInBytes":
//        type: "integer"
//    "storageUsageInBytesSetAt":
//        type: "string"
//        format: "date-time"
//    "updatedAt":
//        type: "string"
//        format: "date-time"
//    "version":
//        type: "string"
//required:
//    - "id"
//    - "projectId"
//    - "version"
//    - "hostname"
//    - "port"
//    - "description"
//    - "createdAt"
//    - "updatedAt"
//    - "name"
//    - "storageUsageInBytes"
//    - "storageUsageInBytesSetAt"
//    - "status"
//    - "statusSetAt"
//example: {"configuration": {"additionalFlags": ["--tcp-keepalive", "300"], "maxMemory": "16Mi", "maxMemoryPolicy": "allkeys-lru", "persistent": "true"}, "createdAt": "2023-03-28T13:15:00.000Z", "description": "My first RedisDatabase!", "finalizers": ["app:appinstallation:28471edf-d266-4d79-8ca8-169e330746db"], "hostname": "redis-xxxxxx.pg-example.db.example.com", "id": "fcfe9aea-84d5-46eb-ac6f-a0ccffa908b1", "name": "redis_xxxxxx", "port": 6379, "projectId": "9f2bddf1-dea6-4441-b4fe-a22ff39caff8", "status": "error", "statusSetAt": "2024-03-05T9:26:32.000Z", "storageUsageInBytes": 10485760, "storageUsageInBytesSetAt": "2024-03-05T9:26:32.000Z", "updatedAt": "2023-03-29T15:50:10.000Z", "version": "7.0"}

//
type RedisDatabase struct {
	Configuration            *RedisDatabaseConfiguration `json:"configuration,omitempty"`
	CreatedAt                string                      `json:"createdAt"`
	Description              string                      `json:"description"`
	Finalizers               []string                    `json:"finalizers,omitempty"`
	Hostname                 string                      `json:"hostname"`
	Id                       uuid.UUID                   `json:"id"`
	Name                     string                      `json:"name"`
	Port                     int64                       `json:"port"`
	ProjectId                uuid.UUID                   `json:"projectId"`
	Status                   DatabaseStatus              `json:"status"`
	StatusSetAt              string                      `json:"statusSetAt"`
	StorageUsageInBytes      int64                       `json:"storageUsageInBytes"`
	StorageUsageInBytesSetAt string                      `json:"storageUsageInBytesSetAt"`
	UpdatedAt                string                      `json:"updatedAt"`
	Version                  string                      `json:"version"`
}

func (o *RedisDatabase) Validate() error {
	if err := func() error {
		if o.Configuration == nil {
			return nil
		}
		return o.Configuration.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property configuration: %w", err)
	}
	if err := func() error {
		if o.Finalizers == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property finalizers: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	return nil
}
