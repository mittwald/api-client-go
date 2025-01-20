package relocation

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "default"
//    - "additionalCompare"

type CreateRelocationRequestBodyAdditionalServicesDataCompare string

const CreateRelocationRequestBodyAdditionalServicesDataCompareDefault CreateRelocationRequestBodyAdditionalServicesDataCompare = "default"
const CreateRelocationRequestBodyAdditionalServicesDataCompareAdditionalCompare CreateRelocationRequestBodyAdditionalServicesDataCompare = "additionalCompare"

func (e CreateRelocationRequestBodyAdditionalServicesDataCompare) Validate() error {
	if e == CreateRelocationRequestBodyAdditionalServicesDataCompareDefault || e == CreateRelocationRequestBodyAdditionalServicesDataCompareAdditionalCompare {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}