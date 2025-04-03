package userclientv2

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "person": {"$ref": "#/components/schemas/de.mittwald.v1.commons.Person"}
// required:
//    - "person"
// description: DeprecatedUpdateAccountRequestBody models the JSON body of a 'deprecated-user-update-account' request

// DeprecatedUpdateAccountRequestBody models the JSON body of a 'deprecated-user-update-account' request
type DeprecatedUpdateAccountRequestBody struct {
	Person commonsv2.Person `json:"person"`
}

func (o *DeprecatedUpdateAccountRequestBody) Validate() error {
	if err := o.Person.Validate(); err != nil {
		return fmt.Errorf("invalid property person: %w", err)
	}
	return nil
}
