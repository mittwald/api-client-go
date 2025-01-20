package backupv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "downloadURL":
//        type: "string"
//        format: "url"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//    "format":
//        type: "string"
//        example: "tar"
//    "phase":
//        type: "string"
//        enum:
//            - ""
//            - "Pending"
//            - "Exporting"
//            - "Failed"
//            - "Completed"
//            - "Expired"
//        example: "Completed"
//    "withPassword":
//        type: "boolean"
//required:
//    - "format"
//    - "withPassword"

//
type ProjectBackupExport struct {
	DownloadURL  *string                   `json:"downloadURL,omitempty"`
	ExpiresAt    *time.Time                `json:"expiresAt,omitempty"`
	Format       string                    `json:"format"`
	Phase        *ProjectBackupExportPhase `json:"phase,omitempty"`
	WithPassword bool                      `json:"withPassword"`
}

func (o *ProjectBackupExport) Validate() error {
	if err := func() error {
		if o.Phase == nil {
			return nil
		}
		return o.Phase.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property phase: %w", err)
	}
	return nil
}
