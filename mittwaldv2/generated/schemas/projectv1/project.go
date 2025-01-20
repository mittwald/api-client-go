package projectv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "backupStorageUsageInBytes":
//        type: "integer"
//        format: "int64"
//    "backupStorageUsageInBytesSetAt":
//        type: "string"
//        format: "date-time"
//    "clusterDomain":
//        type: "string"
//        format: "hostname"
//        example: "project.host"
//    "clusterID":
//        type: "string"
//        description: "deprecated by property clusterId"
//        example: "espelkamp"
//        deprecated: true
//    "clusterId":
//        type: "string"
//        example: "espelkamp"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "customerId":
//        type: "string"
//        example: "f282f1a8-2b15-4b08-9850-6788e3b20136"
//    "description":
//        type: "string"
//        example: "My first Project!"
//    "directories":
//        type: "object"
//        additionalProperties:
//            type: "string"
//        example: {"Home": "/home/p-4e7tz3"}
//    "disableReason": {"$ref": "#/components/schemas/de.mittwald.v1.project.DisableReason"}
//    "disabledAt":
//        type: "string"
//        format: "date-time"
//    "enabled":
//        type: "boolean"
//    "id":
//        type: "string"
//        format: "uuid"
//    "imageRefId":
//        type: "string"
//        format: "uuid"
//    "isReady":
//        type: "boolean"
//        description: "deprecated by property status"
//        deprecated: true
//    "projectHostingId":
//        type: "string"
//        format: "uuid"
//    "readiness": {"$ref": "#/components/schemas/de.mittwald.v1.project.DeprecatedProjectReadinessStatus"}
//    "serverId":
//        type: "string"
//        format: "uuid"
//    "serverShortId":
//        type: "string"
//    "shortId":
//        type: "string"
//        example: "s-4e7tz3"
//    "spec":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.project.VisitorSpec"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.project.HardwareSpec"}
//    "statisticsBaseDomain":
//        type: "string"
//        format: "hostname"
//        example: "pe-prod.staging.mcloud.services"
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.project.ProjectStatus"}
//    "statusSetAt":
//        type: "string"
//        format: "date-time"
//    "webStorageUsageInBytes":
//        type: "integer"
//        format: "int64"
//    "webStorageUsageInBytesSetAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "id"
//    - "shortId"
//    - "description"
//    - "enabled"
//    - "customerId"
//    - "directories"
//    - "createdAt"
//    - "isReady"
//    - "readiness"
//    - "status"
//    - "statusSetAt"
//    - "webStorageUsageInBytes"
//    - "webStorageUsageInBytesSetAt"
//    - "backupStorageUsageInBytes"
//    - "backupStorageUsageInBytesSetAt"

//
type Project struct {
	BackupStorageUsageInBytes      int64                            `json:"backupStorageUsageInBytes"`
	BackupStorageUsageInBytesSetAt time.Time                        `json:"backupStorageUsageInBytesSetAt"`
	ClusterDomain                  *string                          `json:"clusterDomain,omitempty"`
	ClusterID                      *string                          `json:"clusterID,omitempty"`
	ClusterId                      *string                          `json:"clusterId,omitempty"`
	CreatedAt                      time.Time                        `json:"createdAt"`
	CustomerId                     string                           `json:"customerId"`
	Description                    string                           `json:"description"`
	Directories                    map[string]string                `json:"directories"`
	DisableReason                  *DisableReason                   `json:"disableReason,omitempty"`
	DisabledAt                     *time.Time                       `json:"disabledAt,omitempty"`
	Enabled                        bool                             `json:"enabled"`
	Id                             uuid.UUID                        `json:"id"`
	ImageRefId                     *uuid.UUID                       `json:"imageRefId,omitempty"`
	IsReady                        bool                             `json:"isReady"`
	ProjectHostingId               *uuid.UUID                       `json:"projectHostingId,omitempty"`
	Readiness                      DeprecatedProjectReadinessStatus `json:"readiness"`
	ServerId                       *uuid.UUID                       `json:"serverId,omitempty"`
	ServerShortId                  *string                          `json:"serverShortId,omitempty"`
	ShortId                        string                           `json:"shortId"`
	Spec                           *ProjectSpec                     `json:"spec,omitempty"`
	StatisticsBaseDomain           *string                          `json:"statisticsBaseDomain,omitempty"`
	Status                         ProjectStatus                    `json:"status"`
	StatusSetAt                    time.Time                        `json:"statusSetAt"`
	WebStorageUsageInBytes         int64                            `json:"webStorageUsageInBytes"`
	WebStorageUsageInBytesSetAt    time.Time                        `json:"webStorageUsageInBytesSetAt"`
}

func (o *Project) Validate() error {
	if o.Directories == nil {
		return errors.New("property directories is required, but not set")
	}
	if err := func() error {
		if o.DisableReason == nil {
			return nil
		}
		return o.DisableReason.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property disableReason: %w", err)
	}
	if err := o.Readiness.Validate(); err != nil {
		return fmt.Errorf("invalid property readiness: %w", err)
	}
	if err := func() error {
		if o.Spec == nil {
			return nil
		}
		return o.Spec.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property spec: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	return nil
}