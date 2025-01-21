package backupv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "databases":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.backup.DatabaseReference"}
//    "files":
//        type: "boolean"
//required:
//    - "files"

type IgnoredSources struct {
	Databases []DatabaseReference `json:"databases,omitempty"`
	Files     bool                `json:"files"`
}

func (o *IgnoredSources) Validate() error {
	if err := func() error {
		if o.Databases == nil {
			return nil
		}
		return func() error {
			for i := range o.Databases {
				if err := o.Databases[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property databases: %w", err)
	}
	return nil
}
