package storagespacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "childStatistics":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.Statistics"}
//        uniqueItems: true
//    "description":
//        type: "string"
//        example: "My First Project"
//    "id":
//        type: "string"
//        example: "169cea81-2c11-46a4-8f0b-5b0b47caeb78"
//    "kind": {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.StatisticsKind"}
//    "meta": {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.StatisticsMeta"}
//    "name":
//        type: "string"
//        example: "p-zkl8tr"
//    "notificationThresholdInBytes":
//        type: "integer"
//        example: 10000
//    "statisticCategories":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.StatisticsCategory"}
//        uniqueItems: true
//required:
//    - "id"
//    - "kind"
//    - "name"
//    - "meta"

type Statistics struct {
	ChildStatistics              []Statistics         `json:"childStatistics,omitempty"`
	Description                  *string              `json:"description,omitempty"`
	Id                           string               `json:"id"`
	Kind                         StatisticsKind       `json:"kind"`
	Meta                         StatisticsMeta       `json:"meta"`
	Name                         string               `json:"name"`
	NotificationThresholdInBytes *int64               `json:"notificationThresholdInBytes,omitempty"`
	StatisticCategories          []StatisticsCategory `json:"statisticCategories,omitempty"`
}

func (o *Statistics) Validate() error {
	if err := func() error {
		if o.ChildStatistics == nil {
			return nil
		}
		return func() error {
			for i := range o.ChildStatistics {
				if err := o.ChildStatistics[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property childStatistics: %w", err)
	}
	if err := o.Kind.Validate(); err != nil {
		return fmt.Errorf("invalid property kind: %w", err)
	}
	if err := o.Meta.Validate(); err != nil {
		return fmt.Errorf("invalid property meta: %w", err)
	}
	if err := func() error {
		if o.StatisticCategories == nil {
			return nil
		}
		return func() error {
			for i := range o.StatisticCategories {
				if err := o.StatisticCategories[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property statisticCategories: %w", err)
	}
	return nil
}
