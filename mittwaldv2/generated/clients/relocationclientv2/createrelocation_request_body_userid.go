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
// oneOf:
//    - type: "string"
//      minLength: 1
//      format: "uuid"
//    - type: "string"

type CreateRelocationRequestBodyUserID struct {
	AlternativeCreateRelocationRequestBodyUserIDAlternative1 *string
	AlternativeCreateRelocationRequestBodyUserIDAlternative2 *string
}

func (a *CreateRelocationRequestBodyUserID) MarshalJSON() ([]byte, error) {
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative1 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyUserIDAlternative1)
	}
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative2 != nil {
		return json.Marshal(a.AlternativeCreateRelocationRequestBodyUserIDAlternative2)
	}
	return []byte("null"), nil
}

func (a *CreateRelocationRequestBodyUserID) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyUserIDAlternative1 string
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyUserIDAlternative1); err == nil {
		a.AlternativeCreateRelocationRequestBodyUserIDAlternative1 = &alternativeCreateRelocationRequestBodyUserIDAlternative1
		decodedAtLeastOnce = true
	}

	reader.Reset(input)
	var alternativeCreateRelocationRequestBodyUserIDAlternative2 string
	if err := dec.Decode(&alternativeCreateRelocationRequestBodyUserIDAlternative2); err == nil {
		a.AlternativeCreateRelocationRequestBodyUserIDAlternative2 = &alternativeCreateRelocationRequestBodyUserIDAlternative2
		decodedAtLeastOnce = true
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *CreateRelocationRequestBodyUserID) Validate() error {
	// The AlternativeCreateRelocationRequestBodyUserIDAlternative1 subtype does not implement validation, so we consider being non-nil as valid
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative1 != nil {
		return nil
	}
	// The AlternativeCreateRelocationRequestBodyUserIDAlternative2 subtype does not implement validation, so we consider being non-nil as valid
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative2 != nil {
		return nil
	}
	return errors.New("no alternative set")
}

func (a CreateRelocationRequestBodyUserID) String() string {
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative1 != nil {
		return *a.AlternativeCreateRelocationRequestBodyUserIDAlternative1
	}
	if a.AlternativeCreateRelocationRequestBodyUserIDAlternative2 != nil {
		return *a.AlternativeCreateRelocationRequestBodyUserIDAlternative2
	}
	return "null"
}
