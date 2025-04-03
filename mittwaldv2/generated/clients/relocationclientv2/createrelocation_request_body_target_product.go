package relocationclientv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// oneOf:
//    - type: "string"
//      minLength: 1
//    - type: "string"
//      enum:
//        - "Space-Server"
//        - "proSpace"
//        - "Agentur-Server"
//        - "CMS-Hosting"
//        - "Shop-Hosting"
// description: "Help our customer service finding your target account"

type CreateRelocationRequestBodyTargetProduct struct {
	AlternativeCreateRelocationRequestBodyTargetProductAlternative1 *string
	AlternativeCreateRelocationRequestBodyTargetProductAlternative2 *CreateRelocationRequestBodyTargetProductAlternative2
}

func (a *CreateRelocationRequestBodyTargetProduct) MarshalJSON() ([]byte, error) {
	if a.AlternativeCreateRelocationRequestBodyTargetProductAlternative1 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyTargetProductAlternative1)
	}
	if a.AlternativeCreateRelocationRequestBodyTargetProductAlternative2 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyTargetProductAlternative2)
	}
	return []byte("null"), nil
}

func (a *CreateRelocationRequestBodyTargetProduct) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyTargetProductAlternative1 string
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyTargetProductAlternative1); err == nil {
		a.AlternativeCreateRelocationRequestBodyTargetProductAlternative1 = &alternativeCreateRelocationRequestBodyTargetProductAlternative1
		decodedAtLeastOnce = true
	}

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyTargetProductAlternative2 CreateRelocationRequestBodyTargetProductAlternative2
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyTargetProductAlternative2); err == nil {
		//subtype: *generator.StringEnumType
		if vErr := alternativeCreateRelocationRequestBodyTargetProductAlternative2.Validate(); vErr == nil {
			a.AlternativeCreateRelocationRequestBodyTargetProductAlternative2 = &alternativeCreateRelocationRequestBodyTargetProductAlternative2
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *CreateRelocationRequestBodyTargetProduct) Validate() error {
	// The AlternativeCreateRelocationRequestBodyTargetProductAlternative1 subtype does not implement validation, so we consider being non-nil as valid
	if a.AlternativeCreateRelocationRequestBodyTargetProductAlternative1 != nil {
		return nil
	}
	if a.AlternativeCreateRelocationRequestBodyTargetProductAlternative2 != nil {
		return a.AlternativeCreateRelocationRequestBodyTargetProductAlternative2.Validate()
	}
	return errors.New("no alternative set")
}
