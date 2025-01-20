package screenshotv1

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
//    "executedAt":
//        type: "string"
//        format: "date-time"
//    "fileReference":
//        type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "priority":
//        type: "number"
//    "settings": {"$ref": "#/components/schemas/de.mittwald.v1.screenshot.ScreenshotSettings"}
//    "target": {"$ref": "#/components/schemas/de.mittwald.v1.screenshot.Target"}
//    "taskState":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.screenshot.LifecycleState"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.commons.Error"}
//required:
//    - "id"
//    - "settings"
//    - "target"
//    - "priority"

//
type Task struct {
	ExecutedAt    *time.Time         `json:"executedAt,omitempty"`
	FileReference *string            `json:"fileReference,omitempty"`
	Id            uuid.UUID          `json:"id"`
	Priority      float64            `json:"priority"`
	Settings      ScreenshotSettings `json:"settings"`
	Target        Target             `json:"target"`
	TaskState     *TaskTaskState     `json:"taskState,omitempty"`
}

func (o *Task) Validate() error {
	if err := o.Settings.Validate(); err != nil {
		return fmt.Errorf("invalid property settings: %w", err)
	}
	if err := o.Target.Validate(); err != nil {
		return fmt.Errorf("invalid property target: %w", err)
	}
	if err := func() error {
		if o.TaskState == nil {
			return nil
		}
		return o.TaskState.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property taskState: %w", err)
	}
	return nil
}
