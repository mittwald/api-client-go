package contractv1

import (
	"errors"
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "newArticles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.contract.Article"}
//    "scheduledAtDate":
//        type: "string"
//        format: "date-time"
//    "scheduledByUserId":
//        type: "string"
//    "targetDate":
//        type: "string"
//        format: "date-time"
// required:
//    - "scheduledAtDate"
//    - "targetDate"
//    - "newArticles"

type TariffChange struct {
	NewArticles       []Article `json:"newArticles"`
	ScheduledAtDate   time.Time `json:"scheduledAtDate"`
	ScheduledByUserId *string   `json:"scheduledByUserId,omitempty"`
	TargetDate        time.Time `json:"targetDate"`
}

func (o *TariffChange) Validate() error {
	if o.NewArticles == nil {
		return errors.New("property newArticles is required, but not set")
	}
	if err := func() error {
		for i := range o.NewArticles {
			if err := o.NewArticles[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property newArticles: %w", err)
	}
	return nil
}
