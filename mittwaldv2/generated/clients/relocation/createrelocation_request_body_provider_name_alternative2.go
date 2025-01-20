package relocation

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "1und1"
//    - "strato"

type CreateRelocationRequestBodyProviderNameAlternative2 string

const CreateRelocationRequestBodyProviderNameAlternative21Und1 CreateRelocationRequestBodyProviderNameAlternative2 = "1und1"
const CreateRelocationRequestBodyProviderNameAlternative2Strato CreateRelocationRequestBodyProviderNameAlternative2 = "strato"

func (e CreateRelocationRequestBodyProviderNameAlternative2) Validate() error {
	if e == CreateRelocationRequestBodyProviderNameAlternative21Und1 || e == CreateRelocationRequestBodyProviderNameAlternative2Strato {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}