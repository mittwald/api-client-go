package projectv1

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
//    - {"$ref": "#/components/schemas/de.mittwald.v1.project.VisitorSpec"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.project.HardwareSpec"}

type ProjectSpec struct {
	AlternativeVisitorSpec  *VisitorSpec
	AlternativeHardwareSpec *HardwareSpec
}

func (a *ProjectSpec) MarshalJSON() ([]byte, error) {
	if a.AlternativeVisitorSpec != nil {
		return json.Marshal(a.AlternativeVisitorSpec)
	}
	if a.AlternativeHardwareSpec != nil {
		return json.Marshal(a.AlternativeHardwareSpec)
	}
	return []byte("null"), nil
}

func (a *ProjectSpec) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeVisitorSpec VisitorSpec
	if err := dec.Decode(&alternativeVisitorSpec); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeVisitorSpec.Validate(); vErr == nil {
			a.AlternativeVisitorSpec = &alternativeVisitorSpec
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeHardwareSpec HardwareSpec
	if err := dec.Decode(&alternativeHardwareSpec); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeHardwareSpec.Validate(); vErr == nil {
			a.AlternativeHardwareSpec = &alternativeHardwareSpec
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *ProjectSpec) Validate() error {
	if a.AlternativeVisitorSpec != nil {
		return a.AlternativeVisitorSpec.Validate()
	}
	if a.AlternativeHardwareSpec != nil {
		return a.AlternativeHardwareSpec.Validate()
	}
	return errors.New("no alternative set")
}
