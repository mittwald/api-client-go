package relocation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - type: "string"
//      minLength: 1
//    - type: "string"
//      enum:
//        - "1und1"
//        - "strato"
//description: "Name of your provider"

type CreateRelocationRequestBodyProviderName struct {
	AlternativeCreateRelocationRequestBodyProviderNameAlternative1 *string
	AlternativeCreateRelocationRequestBodyProviderNameAlternative2 *CreateRelocationRequestBodyProviderNameAlternative2
}

func (a *CreateRelocationRequestBodyProviderName) MarshalJSON() ([]byte, error) {
	if a.AlternativeCreateRelocationRequestBodyProviderNameAlternative1 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyProviderNameAlternative1)
	}
	if a.AlternativeCreateRelocationRequestBodyProviderNameAlternative2 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyProviderNameAlternative2)
	}
	return []byte("null"), nil
}

func (a *CreateRelocationRequestBodyProviderName) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyProviderNameAlternative1 string
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyProviderNameAlternative1); err == nil {
		a.AlternativeCreateRelocationRequestBodyProviderNameAlternative1 = &alternativeCreateRelocationRequestBodyProviderNameAlternative1
		decodedAtLeastOnce = true
	}

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyProviderNameAlternative2 CreateRelocationRequestBodyProviderNameAlternative2
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyProviderNameAlternative2); err == nil {
		//subtype: *generator.StringEnumType
		if vErr := alternativeCreateRelocationRequestBodyProviderNameAlternative2.Validate(); vErr == nil {
			a.AlternativeCreateRelocationRequestBodyProviderNameAlternative2 = &alternativeCreateRelocationRequestBodyProviderNameAlternative2
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *CreateRelocationRequestBodyProviderName) Validate() error {
	// The AlternativeCreateRelocationRequestBodyProviderNameAlternative1 subtype does not implement validation, so we consider being non-nil as valid
	if a.AlternativeCreateRelocationRequestBodyProviderNameAlternative1 != nil {
		return nil
	}
	if a.AlternativeCreateRelocationRequestBodyProviderNameAlternative2 != nil {
		return a.AlternativeCreateRelocationRequestBodyProviderNameAlternative2.Validate()
	}
	return errors.New("no alternative set")
}
