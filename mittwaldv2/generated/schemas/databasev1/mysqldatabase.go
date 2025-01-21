package databasev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "characterSettings": {"$ref": "#/components/schemas/de.mittwald.v1.database.characterSettings"}
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "description":
//        type: "string"
//    "externalHostname":
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
//    "isReady":
//        type: "boolean"
//    "isShared":
//        type: "boolean"
//    "name":
//        type: "string"
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
//    - "createdAt"
//    - "updatedAt"
//    - "projectId"
//    - "name"
//    - "description"
//    - "version"
//    - "characterSettings"
//    - "hostname"
//    - "isShared"
//    - "isReady"
//    - "storageUsageInBytes"
//    - "storageUsageInBytesSetAt"
//    - "status"
//    - "statusSetAt"
//    - "externalHostname"

type MySqlDatabase struct {
	CharacterSettings        CharacterSettings `json:"characterSettings"`
	CreatedAt                time.Time         `json:"createdAt"`
	Description              string            `json:"description"`
	ExternalHostname         string            `json:"externalHostname"`
	Finalizers               []string          `json:"finalizers,omitempty"`
	Hostname                 string            `json:"hostname"`
	Id                       uuid.UUID         `json:"id"`
	IsReady                  bool              `json:"isReady"`
	IsShared                 bool              `json:"isShared"`
	Name                     string            `json:"name"`
	ProjectId                uuid.UUID         `json:"projectId"`
	Status                   DatabaseStatus    `json:"status"`
	StatusSetAt              time.Time         `json:"statusSetAt"`
	StorageUsageInBytes      int64             `json:"storageUsageInBytes"`
	StorageUsageInBytesSetAt time.Time         `json:"storageUsageInBytesSetAt"`
	UpdatedAt                time.Time         `json:"updatedAt"`
	Version                  string            `json:"version"`
}

func (o *MySqlDatabase) Validate() error {
	if err := o.CharacterSettings.Validate(); err != nil {
		return fmt.Errorf("invalid property characterSettings: %w", err)
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
