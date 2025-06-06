package projectv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "clusterName":
//        type: "string"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "customerId":
//        type: "string"
//        example: "673c107f-75e1-451c-8eaa-5bf101bd2b2c"
//    "description":
//        type: "string"
//        example: "My first Server!"
//    "disabledReason": {"$ref": "#/components/schemas/de.mittwald.v1.project.ServerDisableReason"}
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
//    "machineType": {"$ref": "#/components/schemas/de.mittwald.v1.project.MachineType"}
//    "readiness": {"$ref": "#/components/schemas/de.mittwald.v1.project.DeprecatedServerReadinessStatus"}
//    "shortId":
//        type: "string"
//        example: "s-4e7tz3"
//    "statisticsBaseDomain":
//        type: "string"
//        format: "hostname"
//        example: "pe-prod.staging.mcloud.services"
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.project.ServerStatus"}
//    "storage":
//        type: "string"
//        example: "50Gi"
// required:
//    - "id"
//    - "shortId"
//    - "customerId"
//    - "machineType"
//    - "storage"
//    - "description"
//    - "isReady"
//    - "readiness"
//    - "createdAt"
//    - "clusterName"
//    - "status"

type Server struct {
	ClusterName          string                          `json:"clusterName"`
	CreatedAt            time.Time                       `json:"createdAt"`
	CustomerId           string                          `json:"customerId"`
	Description          string                          `json:"description"`
	DisabledReason       *ServerDisableReason            `json:"disabledReason,omitempty"`
	Id                   string                          `json:"id"`
	ImageRefId           *string                         `json:"imageRefId,omitempty"`
	IsReady              bool                            `json:"isReady"`
	MachineType          MachineType                     `json:"machineType"`
	Readiness            DeprecatedServerReadinessStatus `json:"readiness"`
	ShortId              string                          `json:"shortId"`
	StatisticsBaseDomain *string                         `json:"statisticsBaseDomain,omitempty"`
	Status               ServerStatus                    `json:"status"`
	Storage              string                          `json:"storage"`
}

func (o *Server) Validate() error {
	if err := func() error {
		if o.DisabledReason == nil {
			return nil
		}
		return o.DisabledReason.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property disabledReason: %w", err)
	}
	if err := o.MachineType.Validate(); err != nil {
		return fmt.Errorf("invalid property machineType: %w", err)
	}
	if err := o.Readiness.Validate(); err != nil {
		return fmt.Errorf("invalid property readiness: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	return nil
}
