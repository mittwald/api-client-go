package mailmigrationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "errors": {"$ref": "#/components/schemas/de.mittwald.v1.mailmigration.PossibleCheckErrors"}

type CheckMigrationIsPossibleErrorResponse struct {
	Errors *PossibleCheckErrors `json:"errors,omitempty"`
}

func (o *CheckMigrationIsPossibleErrorResponse) Validate() error {
	if err := func() error {
		if o.Errors == nil {
			return nil
		}
		return o.Errors.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property errors: %w", err)
	}
	return nil
}
