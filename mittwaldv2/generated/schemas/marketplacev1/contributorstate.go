package marketplacev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "enabled"
//    - "disabled"

type ContributorState string

const ContributorStateEnabled ContributorState = "enabled"
const ContributorStateDisabled ContributorState = "disabled"

func (e ContributorState) Validate() error {
	if e == ContributorStateEnabled || e == ContributorStateDisabled {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}