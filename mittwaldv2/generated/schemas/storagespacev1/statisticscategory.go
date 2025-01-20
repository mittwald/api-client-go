package storagespacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "kind": {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.StatisticsCategoryKind"}
//    "resources":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.storagespace.StatisticsResource"}
//        uniqueItems: true
//    "totalUsageInBytes":
//        type: "integer"
//        example: 1000
//required:
//    - "kind"
//    - "totalUsageInBytes"

//
type StatisticsCategory struct {
	Kind              StatisticsCategoryKind `json:"kind"`
	Resources         []StatisticsResource   `json:"resources,omitempty"`
	TotalUsageInBytes int64                  `json:"totalUsageInBytes"`
}

func (o *StatisticsCategory) Validate() error {
	if err := o.Kind.Validate(); err != nil {
		return fmt.Errorf("invalid property kind: %w", err)
	}
	if err := func() error {
		if o.Resources == nil {
			return nil
		}
		return func() error {
			for i := range o.Resources {
				if err := o.Resources[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property resources: %w", err)
	}
	return nil
}